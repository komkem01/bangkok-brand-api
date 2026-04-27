package review

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ProductReview, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "review.Info")
	defer span.End()
	item, err := s.db.GetReviewByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("review.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
