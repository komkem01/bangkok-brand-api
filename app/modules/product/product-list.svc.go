package product

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Product, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "product.List")
	defer span.End()

	items, err := s.db.ListProducts(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("product.list.error: %v", err)
		return nil, err
	}

	log.Infof("product.list.ok count=%d", len(items))
	return items, nil
}
