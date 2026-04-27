package view

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ProductView, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "view.Info")
	defer span.End()
	item, err := s.db.GetProductViewByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("view.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
