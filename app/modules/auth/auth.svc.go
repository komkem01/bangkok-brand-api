package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/hashing"
	"bangkok-brand/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrEmailExists        = errors.New("email already exists")
	ErrMemberNotActive    = errors.New("member is not active")
)

const (
	ContextKeyMemberID = "auth.member_id"
	ContextKeyMember   = "auth.member"
)

type Config struct {
	AccessTokenTTLMinutes int    `conf:""`
	RefreshTokenTTLHours  int    `conf:""`
	TokenIssuer           string `conf:""`
	TokenSecret           string `conf:""`
}

type Options struct {
	Config *config.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.MemberEntity
}

type Service struct {
	*Options
}

type RegisterInput struct {
	Email     string
	Password  string
	Phone     *string
	Firstname *string
	Lastname  *string
}

type LoginInput struct {
	Email    string
	Password string
}

type TokenPair struct {
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	TokenType             string    `json:"token_type"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

type AuthMember struct {
	ID          uuid.UUID        `json:"id"`
	Email       *string          `json:"email"`
	Phone       *string          `json:"phone"`
	Firstname   *string          `json:"firstname_th"`
	Lastname    *string          `json:"lastname_th"`
	Role        ent.MemberRole   `json:"role"`
	Status      ent.MemberStatus `json:"status"`
	IsVerified  bool             `json:"is_verified"`
	LastedLogin *time.Time       `json:"lasted_login"`
}

type AuthResult struct {
	Member AuthMember `json:"member"`
	Tokens TokenPair  `json:"tokens"`
}

type tokenPayload struct {
	Subject string `json:"sub"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Type    string `json:"type"`
	jwt.RegisteredClaims
}

func newService(opts *Options) *Service {
	return &Service{Options: opts}
}

func (s *Service) Register(ctx context.Context, in *RegisterInput) (*AuthResult, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "auth.Register")
	defer span.End()

	email := normalizeEmail(in.Email)
	password := strings.TrimSpace(in.Password)
	if email == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	if len(password) < 8 {
		return nil, ErrInvalidCredentials
	}

	_, err := s.db.GetMemberByEmail(ctx, email)
	if err == nil {
		return nil, ErrEmailExists
	}
	if !errors.Is(err, sql.ErrNoRows) {
		span.RecordError(err)
		log.Errf("auth.register.lookup.error email=%s err=%v", email, err)
		return nil, err
	}

	hashed, err := hashing.HashPassword(password)
	if err != nil {
		span.RecordError(err)
		return nil, fmt.Errorf("hash password: %w", err)
	}

	now := time.Now()
	passwordHash := string(hashed)

	member := &ent.Member{
		Email:       &email,
		Password:    &passwordHash,
		Phone:       in.Phone,
		FirstnameTh: in.Firstname,
		LastnameTh:  in.Lastname,
		Role:        ent.MemberRoleCustomer,
		Status:      ent.MemberStatusActive,
		RegisterdAt: &now,
	}

	created, err := s.db.CreateMember(ctx, member)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil, ErrEmailExists
		}
		span.RecordError(err)
		log.Errf("auth.register.create.error email=%s err=%v", email, err)
		return nil, err
	}

	pair, err := s.newTokenPair(created, now)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	log.Infof("auth.register.ok member_id=%s", created.ID)
	return &AuthResult{
		Member: toAuthMember(created),
		Tokens: *pair,
	}, nil
}

func (s *Service) Login(ctx context.Context, in *LoginInput) (*AuthResult, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "auth.Login")
	defer span.End()

	email := normalizeEmail(in.Email)
	password := strings.TrimSpace(in.Password)
	if email == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	member, err := s.db.GetMemberByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		span.RecordError(err)
		return nil, err
	}

	if member.Status != ent.MemberStatusActive {
		return nil, ErrMemberNotActive
	}

	if member.Password == nil || !hashing.CheckPasswordHash([]byte(*member.Password), []byte(password)) {
		return nil, ErrInvalidCredentials
	}

	now := time.Now()
	if err := s.db.UpdateMemberLastLoginByID(ctx, member.ID, &now); err != nil {
		log.Errf("auth.login.update_last_login.error member_id=%s err=%v", member.ID, err)
	}
	member.LastedLogin = &now

	pair, err := s.newTokenPair(member, now)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	log.Infof("auth.login.ok member_id=%s", member.ID)
	return &AuthResult{
		Member: toAuthMember(member),
		Tokens: *pair,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error) {
	ctx, span, _ := utils.NewLogSpan(ctx, s.tracer, "auth.RefreshToken")
	defer span.End()

	payload, err := s.parseToken(refreshToken, "refresh")
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, ErrInvalidToken
	}

	memberID, err := uuid.Parse(payload.Subject)
	if err != nil {
		return nil, ErrInvalidToken
	}

	member, err := s.db.GetMemberByID(ctx, memberID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	if member.Status != ent.MemberStatusActive {
		return nil, ErrMemberNotActive
	}

	return s.newTokenPair(member, time.Now())
}

func (s *Service) VerifyAccessToken(ctx context.Context, accessToken string) (*ent.Member, error) {
	payload, err := s.parseToken(accessToken, "access")
	if err != nil {
		return nil, ErrInvalidToken
	}

	memberID, err := uuid.Parse(payload.Subject)
	if err != nil {
		return nil, ErrInvalidToken
	}

	member, err := s.db.GetMemberByID(ctx, memberID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	if member.Status != ent.MemberStatusActive {
		return nil, ErrMemberNotActive
	}

	return member, nil
}

func (s *Service) Logout(ctx context.Context) error {
	_, span, _ := utils.NewLogSpan(ctx, s.tracer, "auth.Logout")
	defer span.End()
	return nil
}

func (s *Service) newTokenPair(member *ent.Member, now time.Time) (*TokenPair, error) {
	accessTTL := s.accessTTL()
	refreshTTL := s.refreshTTL()

	accessExp := now.Add(accessTTL)
	refreshExp := now.Add(refreshTTL)

	accessToken, err := s.signToken(tokenPayload{
		Subject: member.ID.String(),
		Email:   memberEmail(member),
		Role:    string(member.Role),
		Type:    "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer(),
			Subject:   member.ID.String(),
			ExpiresAt: jwt.NewNumericDate(accessExp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.signToken(tokenPayload{
		Subject: member.ID.String(),
		Email:   memberEmail(member),
		Role:    string(member.Role),
		Type:    "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer(),
			Subject:   member.ID.String(),
			ExpiresAt: jwt.NewNumericDate(refreshExp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		TokenType:             "Bearer",
		AccessTokenExpiresAt:  accessExp,
		RefreshTokenExpiresAt: refreshExp,
	}, nil
}

func (s *Service) signToken(payload tokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(s.tokenSecret())
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return tokenString, nil
}

func (s *Service) parseToken(token string, tokenType string) (*tokenPayload, error) {
	var payload tokenPayload
	parsed, err := jwt.ParseWithClaims(token, &payload, func(t *jwt.Token) (any, error) {
		if t.Method == nil || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, ErrInvalidToken
		}
		return s.tokenSecret(), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, ErrInvalidToken
	}

	if parsed == nil || !parsed.Valid {
		return nil, ErrInvalidToken
	}

	if payload.Type != tokenType {
		return nil, ErrInvalidToken
	}

	return &payload, nil
}

func (s *Service) accessTTL() time.Duration {
	if s.Config.Val != nil && s.Config.Val.AccessTokenTTLMinutes > 0 {
		return time.Duration(s.Config.Val.AccessTokenTTLMinutes) * time.Minute
	}
	return 15 * time.Minute
}

func (s *Service) refreshTTL() time.Duration {
	if s.Config.Val != nil && s.Config.Val.RefreshTokenTTLHours > 0 {
		return time.Duration(s.Config.Val.RefreshTokenTTLHours) * time.Hour
	}
	return 24 * time.Hour
}

func (s *Service) issuer() string {
	if s.Config.Val != nil && strings.TrimSpace(s.Config.Val.TokenIssuer) != "" {
		return s.Config.Val.TokenIssuer
	}
	return s.Config.AppName()
}

func (s *Service) tokenSecret() []byte {
	if s.Config.Val != nil && strings.TrimSpace(s.Config.Val.TokenSecret) != "" {
		return []byte(s.Config.Val.TokenSecret)
	}
	return []byte(s.Config.AppName() + ":" + s.Config.Environment() + ":" + s.Config.Hostname())
}

func toAuthMember(member *ent.Member) AuthMember {
	return AuthMember{
		ID:          member.ID,
		Email:       member.Email,
		Phone:       member.Phone,
		Firstname:   member.FirstnameTh,
		Lastname:    member.LastnameTh,
		Role:        member.Role,
		Status:      member.Status,
		IsVerified:  member.IsVerified,
		LastedLogin: member.LastedLogin,
	}
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func memberEmail(member *ent.Member) string {
	if member.Email == nil {
		return ""
	}
	return *member.Email
}
