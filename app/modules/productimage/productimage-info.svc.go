package productimage

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productImage.Info")
	defer span.End()

	item, err := s.db.GetProductImageByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("product-image.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("product-image.info.ok id=%s", id)
	return item, nil
}
