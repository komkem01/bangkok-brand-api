package productimage

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductImage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productImage.List")
	defer span.End()

	items, err := s.db.ListProductImages(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-image.list.error: %v", err)
		return nil, err
	}

	log.Infof("product-image.list.ok count=%d", len(items))
	return items, nil
}
