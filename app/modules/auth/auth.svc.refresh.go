package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/codes"
)

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
