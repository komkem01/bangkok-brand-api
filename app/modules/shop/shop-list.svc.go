package shop

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Shop, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shop.List")
	defer span.End()
	items, err := s.db.ListShops(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("shop.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
