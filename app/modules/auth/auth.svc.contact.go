package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"bangkok-brand/app/utils"
)

func (s *Service) ContactInfo(ctx context.Context, accessToken string) (*AuthMember, error) {
	ctx, span, _ := utils.NewLogSpan(ctx, s.tracer, "auth.ContactInfo")
	defer span.End()

	member, err := s.VerifyAccessToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	result := toAuthMember(member)
	return &result, nil
}

func (s *Service) UpdateContactInfo(ctx context.Context, accessToken string, email, phone *string) (*AuthMember, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "auth.UpdateContactInfo")
	defer span.End()

	member, err := s.VerifyAccessToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	updated := *member

	if email != nil {
		normalized := normalizeEmail(*email)
		if normalized == "" {
			return nil, ErrInvalidCredentials
		}
		if member.Email == nil || !strings.EqualFold(*member.Email, normalized) {
			existing, findErr := s.db.GetMemberByEmail(ctx, normalized)
			if findErr == nil && existing.ID != member.ID {
				return nil, ErrEmailExists
			}
			if findErr != nil && !errors.Is(findErr, sql.ErrNoRows) {
				span.RecordError(findErr)
				log.Errf("auth.contact.lookup_email.error email=%s err=%v", normalized, findErr)
				return nil, findErr
			}
		}
		updated.Email = &normalized
	}

	if phone != nil {
		trimmedPhone := strings.TrimSpace(*phone)
		updated.Phone = &trimmedPhone
	}

	result, err := s.db.UpdateMemberByID(ctx, member.ID, &updated)
	if err != nil {
		span.RecordError(err)
		log.Errf("auth.contact.update.error member_id=%s err=%v", member.ID, err)
		return nil, err
	}

	mapped := toAuthMember(result)
	return &mapped, nil
}
