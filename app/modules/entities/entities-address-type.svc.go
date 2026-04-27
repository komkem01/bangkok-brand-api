package entities

import (
	"context"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.AddressTypeEntity = (*Service)(nil)

func (s *Service) ListAddressTypes(ctx context.Context) ([]*ent.AddressType, error) {
	var items []*ent.AddressType
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("name_en ASC").
		OrderExpr("id ASC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetAddressTypeByID(ctx context.Context, id uuid.UUID) (*ent.AddressType, error) {
	item := &ent.AddressType{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) GetAddressTypeByNameEn(ctx context.Context, nameEn string) (*ent.AddressType, error) {
	item := &ent.AddressType{}
	err := s.db.NewSelect().
		Model(item).
		Where("UPPER(name_en) = UPPER(?)", strings.TrimSpace(nameEn)).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateAddressType(ctx context.Context, item *ent.AddressType) (*ent.AddressType, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateAddressTypeByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.AddressType, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.AddressType)(nil)).
		Set("name_th = ?", nameTh).
		Set("name_en = ?", strings.TrimSpace(nameEn)).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetAddressTypeByID(ctx, id)
}

func (s *Service) DeleteAddressTypeByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.AddressType)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
