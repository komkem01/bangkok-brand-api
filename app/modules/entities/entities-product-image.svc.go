package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductImageEntity = (*Service)(nil)

func (s *Service) ListProductImages(ctx context.Context) ([]*ent.ProductImage, error) {
	var items []*ent.ProductImage
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductImageByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error) {
	item := &ent.ProductImage{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductImage(ctx context.Context, p *ent.ProductImage) (*ent.ProductImage, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateProductImageByID(ctx context.Context, id uuid.UUID, p *ent.ProductImage) (*ent.ProductImage, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.ProductImage)(nil)).
		Set("product_id = ?", p.ProductID).
		Set("storage_id = ?", p.StorageID).
		Set("is_main = ?", p.IsMain).
		Set("sort_order = ?", p.SortOrder).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductImageByID(ctx, id)
}

func (s *Service) DeleteProductImageByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductImage)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
