package view

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductView) (*ent.ProductView, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "view.Create")
	defer span.End()
	created, err := s.db.CreateProductView(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("view.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
