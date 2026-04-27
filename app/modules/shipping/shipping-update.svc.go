package shipping

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.ShippingZone) (*ent.ShippingZone, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shipping.Update")
	defer span.End()
	updated, err := s.db.UpdateShippingZoneByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("shipping.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
