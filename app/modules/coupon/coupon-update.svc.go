package coupon

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.Coupon) (*ent.Coupon, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "coupon.Update")
	defer span.End()
	updated, err := s.db.UpdateCouponByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("coupon.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
