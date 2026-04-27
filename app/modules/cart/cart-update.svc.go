package cart

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.Cart) (*ent.Cart, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cart.Update")
	defer span.End()

	updated, err := s.db.UpdateCartByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("cart.update.ok id=%s", id)
	return updated, nil
}
