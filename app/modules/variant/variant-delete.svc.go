package variant

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "variant.Delete")
	defer span.End()
	err := s.db.DeleteVariantByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("variant.delete.error: %v", err)
	}
	return err
}
