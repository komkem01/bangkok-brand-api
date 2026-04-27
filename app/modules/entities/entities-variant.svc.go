package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.VariantEntity = (*Service)(nil)

func (s *Service) ListVariants(ctx context.Context) ([]*ent.ProductVariant, error) {
	var items []*ent.ProductVariant
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetVariantByID(ctx context.Context, id uuid.UUID) (*ent.ProductVariant, error) {
	item := &ent.ProductVariant{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateVariant(ctx context.Context, item *ent.ProductVariant) (*ent.ProductVariant, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateVariantByID(ctx context.Context, id uuid.UUID, item *ent.ProductVariant) (*ent.ProductVariant, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetVariantByID(ctx, id)
}

func (s *Service) DeleteVariantByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductVariant)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
