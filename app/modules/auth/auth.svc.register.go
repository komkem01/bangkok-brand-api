package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/hashing"

	"github.com/google/uuid"
)

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
		GenderID:      in.GenderID,
		PrefixID:      in.PrefixID,
		Email:         &email,
		Password:      &passwordHash,
		Displayname:   in.Displayname,
		FirstnameTh:   in.Firstname,
		LastnameTh:    in.Lastname,
		CitizenID:     in.CitizenID,
		Birthdate:     in.Birthdate,
		Phone:         in.Phone,
		Role:          ent.MemberRoleCustomer,
		Status:        ent.MemberStatusActive,
		ProvinceID:    in.ProvinceID,
		DistrictID:    in.DistrictID,
		SubDistrictID: in.SubDistrictID,
		ZipcodeID:     in.ZipcodeID,
		RegisterdAt:   &now,
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

	if err := s.createDefaultMemberContacts(ctx, created, in.Phone); err != nil {
		span.RecordError(err)
		log.Errf("auth.register.create_contacts.error member_id=%s err=%v", created.ID, err)
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

func (s *Service) createDefaultMemberContacts(ctx context.Context, member *ent.Member, phone *string) error {
	if member == nil || member.Email == nil {
		return nil
	}

	emailTypeID, err := s.getContactTypeIDByNameEn(ctx, "EMAIL")
	if err != nil {
		return err
	}

	_, err = s.ctdb.CreateContact(ctx, &ent.MemberContact{
		MemberID:      &member.ID,
		ContactTypeID: emailTypeID,
		Value:         member.Email,
		IsPrimary:     true,
		IsVerified:    member.IsVerified,
	})
	if err != nil {
		return err
	}

	if phone == nil {
		return nil
	}

	trimmedPhone := strings.TrimSpace(*phone)
	if trimmedPhone == "" {
		return nil
	}

	phoneTypeID, err := s.getContactTypeIDByNameEn(ctx, "PHONE")
	if err != nil {
		return err
	}

	_, err = s.ctdb.CreateContact(ctx, &ent.MemberContact{
		MemberID:      &member.ID,
		ContactTypeID: phoneTypeID,
		Value:         &trimmedPhone,
		IsPrimary:     false,
		IsVerified:    false,
	})
	return err
}

func (s *Service) getContactTypeIDByNameEn(ctx context.Context, nameEn string) (*uuid.UUID, error) {
	if s.cttdb == nil {
		return nil, nil
	}

	item, err := s.cttdb.GetContactTypeByNameEn(ctx, nameEn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &item.ID, nil
}
