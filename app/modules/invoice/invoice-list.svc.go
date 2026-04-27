package invoice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Invoice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "invoice.List")
	defer span.End()
	items, err := s.db.ListInvoices(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("invoice.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
