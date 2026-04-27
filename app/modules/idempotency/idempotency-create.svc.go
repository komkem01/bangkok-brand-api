package idempotency

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.IdempotencyKey) (*ent.IdempotencyKey, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "idempotency.Create")
	defer span.End()
	created, err := s.db.CreateIdempotencyKey(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("idempotency.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
