package productattributevalue

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productAttributeValue.Delete")
	defer span.End()

	if err := s.db.DeleteProductAttributeValueByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("product-attribute-value.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("product-attribute-value.delete.ok id=%s", id)
	return nil
}
