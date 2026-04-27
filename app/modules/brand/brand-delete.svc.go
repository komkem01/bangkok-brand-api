package brand

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "brand.Delete")
	defer span.End()

	if err := s.db.DeleteBrandByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("brand.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("brand.delete.ok id=%s", id)
	return nil
}
