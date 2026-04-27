package shipping

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ShippingZone) (*ent.ShippingZone, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shipping.Create")
	defer span.End()
	created, err := s.db.CreateShippingZone(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("shipping.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
