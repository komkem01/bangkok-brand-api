package cart

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Cart) (*ent.Cart, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "cart.Create")
	defer span.End()

	created, err := s.db.CreateCart(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("cart.create.error: %v", err)
		return nil, err
	}

	log.Infof("cart.create.ok id=%s", created.ID)
	return created, nil
}
