package category

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.Category) (*ent.Category, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "category.Update")
	defer span.End()

	updated, err := s.db.UpdateCategoryByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("category.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("category.update.ok id=%s", id)
	return updated, nil
}
