package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductVariantStockEntity = (*Service)(nil)

func (s *Service) ListProductVariantStocks(ctx context.Context) ([]*ent.ProductVariantStock, error) {
	var items []*ent.ProductVariantStock
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("updated_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductVariantStockByID(ctx context.Context, id uuid.UUID) (*ent.ProductVariantStock, error) {
	item := &ent.ProductVariantStock{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductVariantStock(ctx context.Context, item *ent.ProductVariantStock) (*ent.ProductVariantStock, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateProductVariantStockByID(ctx context.Context, id uuid.UUID, item *ent.ProductVariantStock) (*ent.ProductVariantStock, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductVariantStockByID(ctx, id)
}

func (s *Service) DeleteProductVariantStockByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductVariantStock)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
