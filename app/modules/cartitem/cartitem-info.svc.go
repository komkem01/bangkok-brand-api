package cartitem

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.CartItem, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cartItem.Info")
	defer span.End()

	item, err := s.db.GetCartItemByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart-item.info.error id=%s: %v", id, err)
		return nil, err
	}

	return item, nil
}
