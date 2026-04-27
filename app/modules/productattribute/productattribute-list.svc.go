package productattribute

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductAttribute, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttribute.List")
	defer span.End()

	items, err := s.db.ListProductAttributes(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute.list.error: %v", err)
		return nil, err
	}

	log.Infof("product-attribute.list.ok count=%d", len(items))
	return items, nil
}
