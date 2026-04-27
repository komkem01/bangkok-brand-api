package entities

import (
	"context"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ContactTypeEntity = (*Service)(nil)

func (s *Service) ListContactTypes(ctx context.Context) ([]*ent.ContactType, error) {
	var items []*ent.ContactType
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("name_en ASC").
		OrderExpr("id ASC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetContactTypeByID(ctx context.Context, id uuid.UUID) (*ent.ContactType, error) {
	item := &ent.ContactType{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) GetContactTypeByNameEn(ctx context.Context, nameEn string) (*ent.ContactType, error) {
	item := &ent.ContactType{}
	err := s.db.NewSelect().
		Model(item).
		Where("UPPER(name_en) = UPPER(?)", strings.TrimSpace(nameEn)).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateContactType(ctx context.Context, item *ent.ContactType) (*ent.ContactType, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateContactTypeByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.ContactType, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.ContactType)(nil)).
		Set("name_th = ?", nameTh).
		Set("name_en = ?", strings.TrimSpace(nameEn)).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetContactTypeByID(ctx, id)
}

func (s *Service) DeleteContactTypeByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ContactType)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
