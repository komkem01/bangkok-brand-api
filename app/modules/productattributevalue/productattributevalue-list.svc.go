package productattributevalue

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductAttributeValue, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttributeValue.List")
	defer span.End()

	items, err := s.db.ListProductAttributeValues(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute-value.list.error: %v", err)
		return nil, err
	}

	log.Infof("product-attribute-value.list.ok count=%d", len(items))
	return items, nil
}
