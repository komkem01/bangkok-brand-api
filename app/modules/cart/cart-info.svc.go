package cart

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Cart, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cart.Info")
	defer span.End()

	item, err := s.db.GetCartByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart.info.error id=%s: %v", id, err)
		return nil, err
	}

	return item, nil
}
