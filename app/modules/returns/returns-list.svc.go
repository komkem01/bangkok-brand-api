package returns

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ReturnRequest, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "returns.List")
	defer span.End()
	items, err := s.db.ListReturnRequests(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("returns.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
