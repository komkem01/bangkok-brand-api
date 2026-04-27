package product

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Product) (*ent.Product, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "product.Create")
	defer span.End()

	created, err := s.db.CreateProduct(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("product.create.error: %v", err)
		return nil, err
	}

	log.Infof("product.create.ok id=%s", created.ID)
	return created, nil
}
