package category

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "category.Info")
	defer span.End()

	item, err := s.db.GetCategoryByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("category.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("category.info.ok id=%s", id)
	return item, nil
}
