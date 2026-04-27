package review

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductReview, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "review.List")
	defer span.End()
	items, err := s.db.ListReviews(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("review.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
