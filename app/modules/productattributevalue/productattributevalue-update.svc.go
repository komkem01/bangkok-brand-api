package productattributevalue

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.ProductAttributeValue) (*ent.ProductAttributeValue, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttributeValue.Update")
	defer span.End()

	updated, err := s.db.UpdateProductAttributeValueByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute-value.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("product-attribute-value.update.ok id=%s", id)
	return updated, nil
}
