package idempotency

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.IdempotencyKey, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "idempotency.List")
	defer span.End()
	items, err := s.db.ListIdempotencyKeys(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("idempotency.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
