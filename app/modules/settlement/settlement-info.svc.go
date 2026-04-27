package settlement

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.SettlementBatch, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "settlement.Info")
	defer span.End()
	item, err := s.db.GetSettlementByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("settlement.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
