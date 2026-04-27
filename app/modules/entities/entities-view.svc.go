package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductViewEntity = (*Service)(nil)

func (s *Service) ListProductViews(ctx context.Context) ([]*ent.ProductView, error) {
	var items []*ent.ProductView
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("viewed_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductViewByID(ctx context.Context, id uuid.UUID) (*ent.ProductView, error) {
	item := &ent.ProductView{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductView(ctx context.Context, item *ent.ProductView) (*ent.ProductView, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateProductViewByID(ctx context.Context, id uuid.UUID, item *ent.ProductView) (*ent.ProductView, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductViewByID(ctx, id)
}

func (s *Service) DeleteProductViewByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductView)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
