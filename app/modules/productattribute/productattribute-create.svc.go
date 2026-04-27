package productattribute

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductAttribute) (*ent.ProductAttribute, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttribute.Create")
	defer span.End()

	created, err := s.db.CreateProductAttribute(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute.create.error: %v", err)
		return nil, err
	}

	log.Infof("product-attribute.create.ok id=%s", created.ID)
	return created, nil
}
