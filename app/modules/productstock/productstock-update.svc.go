package productstock

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.ProductStock) (*ent.ProductStock, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productStock.Update")
	defer span.End()

	updated, err := s.db.UpdateProductStockByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-stock.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("product-stock.update.ok id=%s", id)
	return updated, nil
}
