package productstock

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productStock.Delete")
	defer span.End()

	if err := s.db.DeleteProductStockByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("product-stock.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("product-stock.delete.ok id=%s", id)
	return nil
}
