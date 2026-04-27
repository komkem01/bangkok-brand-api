package shipping

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ShippingZone, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shipping.Info")
	defer span.End()
	item, err := s.db.GetShippingZoneByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("shipping.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
