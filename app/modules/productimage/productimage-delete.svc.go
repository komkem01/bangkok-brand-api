package productimage

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "productImage.Delete")
	defer span.End()

	if err := s.db.DeleteProductImageByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("product-image.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("product-image.delete.ok id=%s", id)
	return nil
}
