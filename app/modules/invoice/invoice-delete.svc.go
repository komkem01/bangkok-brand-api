package invoice

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "invoice.Delete")
	defer span.End()
	err := s.db.DeleteInvoiceByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("invoice.delete.error: %v", err)
	}
	return err
}
