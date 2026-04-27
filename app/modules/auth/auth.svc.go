package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"
	"bangkok-brand/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	ctdb   entitiesinf.ContactEntity
	cttdb  entitiesinf.ContactTypeEntity
}

type Service struct {
	*Options
}

type RegisterInput struct {
	GenderID      *uuid.UUID
	PrefixID      *uuid.UUID
	Email         string
	Password      string
	Displayname   *string
	Firstname     *string
	Lastname      *string
	CitizenID     *string
	Birthdate     *time.Time
	Phone         *string
	ProvinceID    *uuid.UUID
	DistrictID    *uuid.UUID
	SubDistrictID *uuid.UUID
	ZipcodeID     *uuid.UUID
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
