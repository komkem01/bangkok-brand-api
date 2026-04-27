package returns

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "returns.Delete")
	defer span.End()
	err := s.db.DeleteReturnRequestByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("returns.delete.error: %v", err)
	}
	return err
}
