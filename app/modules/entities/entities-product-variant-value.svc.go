package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductVariantValueEntity = (*Service)(nil)

func (s *Service) ListProductVariantValues(ctx context.Context) ([]*ent.ProductVariantValue, error) {
	var items []*ent.ProductVariantValue
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductVariantValueByID(ctx context.Context, id uuid.UUID) (*ent.ProductVariantValue, error) {
	item := &ent.ProductVariantValue{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductVariantValue(ctx context.Context, item *ent.ProductVariantValue) (*ent.ProductVariantValue, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateProductVariantValueByID(ctx context.Context, id uuid.UUID, item *ent.ProductVariantValue) (*ent.ProductVariantValue, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductVariantValueByID(ctx, id)
}

func (s *Service) DeleteProductVariantValueByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductVariantValue)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
