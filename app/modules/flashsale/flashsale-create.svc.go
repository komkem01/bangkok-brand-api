package flashsale

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.FlashSaleEvent) (*ent.FlashSaleEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "flashsale.Create")
	defer span.End()
	created, err := s.db.CreateFlashSaleEvent(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("flashsale.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
