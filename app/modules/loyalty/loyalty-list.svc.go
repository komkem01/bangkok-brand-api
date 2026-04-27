package loyalty

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.PointTransaction, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "loyalty.List")
	defer span.End()
	items, err := s.db.ListPointTransactions(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("loyalty.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
