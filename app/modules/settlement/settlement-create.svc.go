package settlement

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.SettlementBatch) (*ent.SettlementBatch, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "settlement.Create")
	defer span.End()
	created, err := s.db.CreateSettlement(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("settlement.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
