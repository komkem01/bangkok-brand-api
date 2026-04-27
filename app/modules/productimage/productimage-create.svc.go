package productimage

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductImage) (*ent.ProductImage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productImage.Create")
	defer span.End()

	created, err := s.db.CreateProductImage(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-image.create.error: %v", err)
		return nil, err
	}

	log.Infof("product-image.create.ok id=%s", created.ID)
	return created, nil
}
