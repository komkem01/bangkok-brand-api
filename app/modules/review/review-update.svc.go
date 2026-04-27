package review

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.ProductReview) (*ent.ProductReview, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "review.Update")
	defer span.End()
	updated, err := s.db.UpdateReviewByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("review.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
