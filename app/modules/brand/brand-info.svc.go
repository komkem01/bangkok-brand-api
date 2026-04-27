package brand

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Brand, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "brand.Info")
	defer span.End()

	item, err := s.db.GetBrandByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("brand.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("brand.info.ok id=%s", id)
	return item, nil
}
