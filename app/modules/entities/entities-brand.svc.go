package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.BrandEntity = (*Service)(nil)

func (s *Service) ListBrands(ctx context.Context) ([]*ent.Brand, error) {
	var items []*ent.Brand
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetBrandByID(ctx context.Context, id uuid.UUID) (*ent.Brand, error) {
	item := &ent.Brand{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateBrand(ctx context.Context, b *ent.Brand) (*ent.Brand, error) {
	_, err := s.db.NewInsert().
		Model(b).
		Returning("*").
		Exec(ctx)
	return b, err
}

func (s *Service) UpdateBrandByID(ctx context.Context, id uuid.UUID, b *ent.Brand) (*ent.Brand, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Brand)(nil)).
		Set("name_th = ?", b.NameTh).
		Set("name_en = ?", b.NameEn).
		Set("logo_id = ?", b.LogoID).
		Set("description = ?", b.Description).
		Set("is_active = ?", b.IsActive).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetBrandByID(ctx, id)
}

func (s *Service) DeleteBrandByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Brand)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
