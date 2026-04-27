package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ReviewEntity = (*Service)(nil)

func (s *Service) ListReviews(ctx context.Context) ([]*ent.ProductReview, error) {
	var items []*ent.ProductReview
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetReviewByID(ctx context.Context, id uuid.UUID) (*ent.ProductReview, error) {
	item := &ent.ProductReview{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateReview(ctx context.Context, item *ent.ProductReview) (*ent.ProductReview, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateReviewByID(ctx context.Context, id uuid.UUID, item *ent.ProductReview) (*ent.ProductReview, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetReviewByID(ctx, id)
}

func (s *Service) DeleteReviewByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductReview)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
