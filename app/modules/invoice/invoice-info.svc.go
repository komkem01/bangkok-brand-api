package invoice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Invoice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "invoice.Info")
	defer span.End()
	item, err := s.db.GetInvoiceByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("invoice.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
