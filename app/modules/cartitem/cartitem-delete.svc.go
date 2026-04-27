package cartitem

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cartItem.Delete")
	defer span.End()

	if err := s.db.DeleteCartItemByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("cart-item.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("cart-item.delete.ok id=%s", id)
	return nil
}
