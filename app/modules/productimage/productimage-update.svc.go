package productimage

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.ProductImage) (*ent.ProductImage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productImage.Update")
	defer span.End()

	updated, err := s.db.UpdateProductImageByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-image.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("product-image.update.ok id=%s", id)
	return updated, nil
}
