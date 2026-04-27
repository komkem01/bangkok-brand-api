package loyalty

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.PointTransaction) (*ent.PointTransaction, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "loyalty.Create")
	defer span.End()
	created, err := s.db.CreatePointTransaction(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("loyalty.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
