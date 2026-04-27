package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.AddressEntity = (*Service)(nil)

func (s *Service) ListAddresses(ctx context.Context) ([]*ent.MemberAddress, error) {
	var addresses []*ent.MemberAddress
	err := s.db.NewSelect().
		Model(&addresses).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return addresses, err
}

func (s *Service) ListAddressesByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberAddress, error) {
	var addresses []*ent.MemberAddress
	err := s.db.NewSelect().
		Model(&addresses).
		Where("member_id = ?", memberID).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return addresses, err
}

func (s *Service) GetAddressByID(ctx context.Context, id uuid.UUID) (*ent.MemberAddress, error) {
	address := &ent.MemberAddress{}
	err := s.db.NewSelect().
		Model(address).
		Where("id = ?", id).
		Scan(ctx)
	return address, err
}

func (s *Service) CreateAddress(ctx context.Context, a *ent.MemberAddress) (*ent.MemberAddress, error) {
	_, err := s.db.NewInsert().
		Model(a).
		Returning("*").
		Exec(ctx)
	return a, err
}

func (s *Service) UpdateAddressByID(ctx context.Context, id uuid.UUID, a *ent.MemberAddress) (*ent.MemberAddress, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.MemberAddress)(nil)).
		Set("member_id = ?", a.MemberID).
		Set("address_type_id = ?", a.AddressTypeID).
		Set("address_name = ?", a.AddressName).
		Set("recipient_name = ?", a.RecipientName).
		Set("recipient_phone = ?", a.RecipientPhone).
		Set("address_detail = ?", a.AddressDetail).
		Set("province_id = ?", a.ProvinceID).
		Set("district_id = ?", a.DistrictID).
		Set("sub_district_id = ?", a.SubDistrictID).
		Set("zipcode_id = ?", a.ZipcodeID).
		Set("is_default = ?", a.IsDefault).
		Set("latitude = ?", a.Latitude).
		Set("longitude = ?", a.Longitude).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetAddressByID(ctx, id)
}

func (s *Service) DeleteAddressByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.MemberAddress)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
