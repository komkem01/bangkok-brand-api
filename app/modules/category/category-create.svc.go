package category

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Category) (*ent.Category, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "category.Create")
	defer span.End()

	created, err := s.db.CreateCategory(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("category.create.error: %v", err)
		return nil, err
	}

	log.Infof("category.create.ok id=%s", created.ID)
	return created, nil
}
