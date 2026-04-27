package review

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductReview) (*ent.ProductReview, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "review.Create")
	defer span.End()
	created, err := s.db.CreateReview(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("review.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
