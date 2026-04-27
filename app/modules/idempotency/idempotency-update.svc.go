package idempotency

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.IdempotencyKey) (*ent.IdempotencyKey, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "idempotency.Update")
	defer span.End()
	updated, err := s.db.UpdateIdempotencyKeyByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("idempotency.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
