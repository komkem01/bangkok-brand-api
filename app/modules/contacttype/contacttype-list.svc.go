package contacttype

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ContactType, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contacttype.List")
	defer span.End()

	items, err := s.db.ListContactTypes(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("contacttype.list.error: %v", err)
		return nil, err
	}

	return items, nil
}
