package productattributevalue

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductAttributeValue) (*ent.ProductAttributeValue, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttributeValue.Create")
	defer span.End()

	created, err := s.db.CreateProductAttributeValue(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute-value.create.error: %v", err)
		return nil, err
	}

	log.Infof("product-attribute-value.create.ok id=%s", created.ID)
	return created, nil
}
