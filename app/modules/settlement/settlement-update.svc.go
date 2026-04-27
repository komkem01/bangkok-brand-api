package settlement

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.SettlementBatch) (*ent.SettlementBatch, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "settlement.Update")
	defer span.End()
	updated, err := s.db.UpdateSettlementByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("settlement.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
