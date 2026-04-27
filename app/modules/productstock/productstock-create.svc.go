package productstock

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductStock) (*ent.ProductStock, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productStock.Create")
	defer span.End()

	created, err := s.db.CreateProductStock(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-stock.create.error: %v", err)
		return nil, err
	}

	log.Infof("product-stock.create.ok id=%s", created.ID)
	return created, nil
}
