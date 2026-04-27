package view

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.ProductView) (*ent.ProductView, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "view.Update")
	defer span.End()
	updated, err := s.db.UpdateProductViewByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("view.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
