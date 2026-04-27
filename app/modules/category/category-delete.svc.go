package category

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "category.Delete")
	defer span.End()

	if err := s.db.DeleteCategoryByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("category.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("category.delete.ok id=%s", id)
	return nil
}
