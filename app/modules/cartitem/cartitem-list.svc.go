package cartitem

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.CartItem, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cartItem.List")
	defer span.End()

	items, err := s.db.ListCartItems(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart-item.list.error: %v", err)
		return nil, err
	}

	log.Infof("cart-item.list.ok count=%d", len(items))
	return items, nil
}
