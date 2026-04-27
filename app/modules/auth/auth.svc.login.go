package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/hashing"
)

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
