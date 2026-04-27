package flashsale

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.FlashSaleEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "flashsale.Info")
	defer span.End()
	item, err := s.db.GetFlashSaleEventByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("flashsale.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
