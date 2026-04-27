package cart

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Cart, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cart.List")
	defer span.End()

	items, err := s.db.ListCarts(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart.list.error: %v", err)
		return nil, err
	}

	log.Infof("cart.list.ok count=%d", len(items))
	return items, nil
}
