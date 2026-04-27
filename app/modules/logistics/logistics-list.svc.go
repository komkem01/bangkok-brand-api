package logistics

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.LogisticsProvider, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "logistics.List")
	defer span.End()
	items, err := s.db.ListLogisticsProviders(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("logistics.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
