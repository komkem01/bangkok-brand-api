package loyalty

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.PointTransaction, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "loyalty.Info")
	defer span.End()
	item, err := s.db.GetPointTransactionByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("loyalty.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
