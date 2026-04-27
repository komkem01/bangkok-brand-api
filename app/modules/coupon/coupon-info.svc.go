package coupon

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Coupon, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "coupon.Info")
	defer span.End()
	item, err := s.db.GetCouponByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("coupon.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
