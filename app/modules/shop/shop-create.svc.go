package shop

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Shop) (*ent.Shop, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shop.Create")
	defer span.End()
	created, err := s.db.CreateShop(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("shop.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
