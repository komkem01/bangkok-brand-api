package brand

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.Brand) (*ent.Brand, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "brand.Update")
	defer span.End()

	updated, err := s.db.UpdateBrandByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("brand.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("brand.update.ok id=%s", id)
	return updated, nil
}
