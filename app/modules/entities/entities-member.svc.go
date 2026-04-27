package entities

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.MemberEntity = (*Service)(nil)

const createMemberMaxRetries = 10

func (s *Service) ListMembers(ctx context.Context) ([]*ent.Member, error) {
	var members []*ent.Member
	err := s.db.NewSelect().
		Model(&members).
		Where("deleted_at IS NULL").
		OrderExpr("created_at DESC").
		Scan(ctx)
	return members, err
}

func (s *Service) GetMemberByID(ctx context.Context, id uuid.UUID) (*ent.Member, error) {
	member := &ent.Member{}
	err := s.db.NewSelect().
		Model(member).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	return member, err
}

func (s *Service) GetMemberByEmail(ctx context.Context, email string) (*ent.Member, error) {
	member := &ent.Member{}
	err := s.db.NewSelect().
		Model(member).
		Where("LOWER(email) = LOWER(?)", email).
		Where("deleted_at IS NULL").
		Scan(ctx)
	return member, err
}

func (s *Service) CreateMember(ctx context.Context, member *ent.Member) (*ent.Member, error) {
	for i := 0; i < createMemberMaxRetries; i++ {
		if member.MemberNo == nil || strings.TrimSpace(*member.MemberNo) == "" {
			generated, err := generateMemberNo()
			if err != nil {
				return nil, err
			}
			member.MemberNo = &generated
		}

		_, err := s.db.NewInsert().
			Model(member).
			Exec(ctx)
		if err == nil {
			return s.GetMemberByID(ctx, member.ID)
		}

		if isDuplicateMemberNoError(err) {
			member.MemberNo = nil
			continue
		}

		return nil, err
	}

	return nil, fmt.Errorf("failed to generate unique member_no after %d attempts", createMemberMaxRetries)
}

func generateMemberNo() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%06d", n.Int64()), nil
}

func isDuplicateMemberNoError(err error) bool {
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "duplicate") && strings.Contains(errMsg, "member_no")
}

func (s *Service) UpdateMemberLastLoginByID(ctx context.Context, id uuid.UUID, lastedLogin *time.Time) error {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Member)(nil)).
		Set("lasted_login = ?", lastedLogin).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

func (s *Service) UpdateMemberByID(ctx context.Context, id uuid.UUID, member *ent.Member) (*ent.Member, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Member)(nil)).
		Set("gender_id = ?", member.GenderID).
		Set("prefix_id = ?", member.PrefixID).
		Set("email = ?", member.Email).
		Set("password = ?", member.Password).
		Set("member_no = ?", member.MemberNo).
		Set("profile_image_id = ?", member.ProfileImageID).
		Set("displayname = ?", member.Displayname).
		Set("firstname_th = ?", member.FirstnameTh).
		Set("lastname_th = ?", member.LastnameTh).
		Set("citizen_id = ?", member.CitizenID).
		Set("birthdate = ?", member.Birthdate).
		Set("phone = ?", member.Phone).
		Set("role = ?", member.Role).
		Set("status = ?", member.Status).
		Set("province_id = ?", member.ProvinceID).
		Set("district_id = ?", member.DistrictID).
		Set("sub_district_id = ?", member.SubDistrictID).
		Set("zipcode_id = ?", member.ZipcodeID).
		Set("registerd_at = ?", member.RegisterdAt).
		Set("lasted_login = ?", member.LastedLogin).
		Set("is_verified = ?", member.IsVerified).
		Set("total_points = ?", member.TotalPoints).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetMemberByID(ctx, id)
}

func (s *Service) DeleteMemberByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewUpdate().
		Model((*ent.Member)(nil)).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
