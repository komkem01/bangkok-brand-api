package coupon

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Coupon, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "coupon.List")
	defer span.End()
	items, err := s.db.ListCoupons(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("coupon.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
