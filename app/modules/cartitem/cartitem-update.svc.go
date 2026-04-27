package cartitem

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.CartItem) (*ent.CartItem, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cartItem.Update")
	defer span.End()

	updated, err := s.db.UpdateCartItemByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart-item.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("cart-item.update.ok id=%s", id)
	return updated, nil
}
