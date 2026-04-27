package view

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductView, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "view.List")
	defer span.End()
	items, err := s.db.ListProductViews(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("view.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
