package product

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Product, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "product.Info")
	defer span.End()

	item, err := s.db.GetProductByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("product.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("product.info.ok id=%s", id)
	return item, nil
}
