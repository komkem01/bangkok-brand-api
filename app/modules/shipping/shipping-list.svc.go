package shipping

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ShippingZone, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shipping.List")
	defer span.End()
	items, err := s.db.ListShippingZones(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("shipping.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
