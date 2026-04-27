package contacttype

import (
	"context"
	"strings"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ContactType) (*ent.ContactType, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contacttype.Create")
	defer span.End()

	item.NameEn = strings.ToUpper(strings.TrimSpace(item.NameEn))
	created, err := s.db.CreateContactType(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("contacttype.create.error: %v", err)
		return nil, err
	}

	return created, nil
}
