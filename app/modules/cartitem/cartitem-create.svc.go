package cartitem

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.CartItem) (*ent.CartItem, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cartItem.Create")
	defer span.End()

	created, err := s.db.CreateCartItem(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart-item.create.error: %v", err)
		return nil, err
	}

	log.Infof("cart-item.create.ok id=%s", created.ID)
	return created, nil
}
