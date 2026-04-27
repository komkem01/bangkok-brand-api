package flashsale

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.FlashSaleEvent) (*ent.FlashSaleEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "flashsale.Update")
	defer span.End()
	updated, err := s.db.UpdateFlashSaleEventByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("flashsale.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
