package contacttype

import (
	"context"
	"strings"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	NameTh   string
	NameEn   string
	IsActive bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, in UpdateInput) (*ent.ContactType, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contacttype.Update")
	defer span.End()

	item, err := s.db.UpdateContactTypeByID(ctx, id, strings.TrimSpace(in.NameTh), strings.ToUpper(strings.TrimSpace(in.NameEn)), in.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("contacttype.update.error id=%s: %v", id, err)
		return nil, err
	}

	return item, nil
}
