package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductStockEntity = (*Service)(nil)

func (s *Service) ListProductStocks(ctx context.Context) ([]*ent.ProductStock, error) {
	var items []*ent.ProductStock
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("updated_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductStockByID(ctx context.Context, id uuid.UUID) (*ent.ProductStock, error) {
	item := &ent.ProductStock{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductStock(ctx context.Context, p *ent.ProductStock) (*ent.ProductStock, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateProductStockByID(ctx context.Context, id uuid.UUID, p *ent.ProductStock) (*ent.ProductStock, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.ProductStock)(nil)).
		Set("product_id = ?", p.ProductID).
		Set("quantity = ?", p.Quantity).
		Set("low_stock_threshold = ?", p.LowStockThreshold).
		Set("last_restocked_at = ?", p.LastRestockedAt).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductStockByID(ctx, id)
}

func (s *Service) DeleteProductStockByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductStock)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
