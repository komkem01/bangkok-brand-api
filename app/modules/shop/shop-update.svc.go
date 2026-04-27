package shop

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.Shop) (*ent.Shop, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shop.Update")
	defer span.End()
	updated, err := s.db.UpdateShopByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("shop.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
