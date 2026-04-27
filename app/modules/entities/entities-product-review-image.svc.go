package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductReviewImageEntity = (*Service)(nil)

func (s *Service) ListProductReviewImages(ctx context.Context) ([]*ent.ProductReviewImage, error) {
	var items []*ent.ProductReviewImage
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductReviewImageByID(ctx context.Context, id uuid.UUID) (*ent.ProductReviewImage, error) {
	item := &ent.ProductReviewImage{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductReviewImage(ctx context.Context, item *ent.ProductReviewImage) (*ent.ProductReviewImage, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateProductReviewImageByID(ctx context.Context, id uuid.UUID, item *ent.ProductReviewImage) (*ent.ProductReviewImage, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductReviewImageByID(ctx, id)
}

func (s *Service) DeleteProductReviewImageByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductReviewImage)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
