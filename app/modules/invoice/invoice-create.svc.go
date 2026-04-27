package invoice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Invoice) (*ent.Invoice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "invoice.Create")
	defer span.End()
	created, err := s.db.CreateInvoice(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("invoice.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
