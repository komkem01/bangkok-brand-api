package loyalty

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.PointTransaction) (*ent.PointTransaction, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "loyalty.Update")
	defer span.End()
	updated, err := s.db.UpdatePointTransactionByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("loyalty.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
