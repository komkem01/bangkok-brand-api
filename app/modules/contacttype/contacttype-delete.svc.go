package contacttype

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contacttype.Delete")
	defer span.End()

	if err := s.db.DeleteContactTypeByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("contacttype.delete.error id=%s: %v", id, err)
		return err
	}

	return nil
}
