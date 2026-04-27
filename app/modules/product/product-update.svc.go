package product

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.Product) (*ent.Product, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "product.Update")
	defer span.End()

	updated, err := s.db.UpdateProductByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("product.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("product.update.ok id=%s", id)
	return updated, nil
}
