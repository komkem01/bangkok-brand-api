package idempotency

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.IdempotencyKey, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "idempotency.Info")
	defer span.End()
	item, err := s.db.GetIdempotencyKeyByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("idempotency.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
