package logistics

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.LogisticsProvider, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "logistics.Info")
	defer span.End()
	item, err := s.db.GetLogisticsProviderByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("logistics.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
