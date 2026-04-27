package flashsale

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.FlashSaleEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "flashsale.List")
	defer span.End()
	items, err := s.db.ListFlashSaleEvents(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("flashsale.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
