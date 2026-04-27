package productstock

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductStock, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productStock.List")
	defer span.End()

	items, err := s.db.ListProductStocks(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-stock.list.error: %v", err)
		return nil, err
	}

	log.Infof("product-stock.list.ok count=%d", len(items))
	return items, nil
}
