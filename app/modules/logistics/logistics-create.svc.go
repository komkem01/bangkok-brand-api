package logistics

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.LogisticsProvider) (*ent.LogisticsProvider, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "logistics.Create")
	defer span.End()
	created, err := s.db.CreateLogisticsProvider(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("logistics.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
