package productattribute

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ProductAttribute, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttribute.Info")
	defer span.End()

	item, err := s.db.GetProductAttributeByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-attribute.info.error id=%s: %v", id, err)
		return nil, err
	}

	return item, nil
}
