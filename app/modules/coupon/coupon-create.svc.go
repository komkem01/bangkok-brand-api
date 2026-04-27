package coupon

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Coupon) (*ent.Coupon, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "coupon.Create")
	defer span.End()
	created, err := s.db.CreateCoupon(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("coupon.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
