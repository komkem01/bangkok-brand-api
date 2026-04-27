package logistics

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.LogisticsProvider) (*ent.LogisticsProvider, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "logistics.Update")
	defer span.End()
	updated, err := s.db.UpdateLogisticsProviderByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("logistics.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
