package category

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Category, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "category.List")
	defer span.End()

	items, err := s.db.ListCategories(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("category.list.error: %v", err)
		return nil, err
	}

	log.Infof("category.list.ok count=%d", len(items))
	return items, nil
}
