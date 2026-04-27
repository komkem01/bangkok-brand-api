package productattributevalue

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ProductAttributeValue, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttributeValue.Info")
	defer span.End()

	item, err := s.db.GetProductAttributeValueByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute-value.info.error id=%s: %v", id, err)
		return nil, err
	}

	return item, nil
}
