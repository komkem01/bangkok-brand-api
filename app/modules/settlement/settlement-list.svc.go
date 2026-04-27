package settlement

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.SettlementBatch, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "settlement.List")
	defer span.End()
	items, err := s.db.ListSettlements(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("settlement.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
