package invoice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.Invoice) (*ent.Invoice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "invoice.Update")
	defer span.End()
	updated, err := s.db.UpdateInvoiceByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("invoice.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
