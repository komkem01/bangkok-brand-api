package contacttype

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ContactType, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contacttype.Info")
	defer span.End()

	item, err := s.db.GetContactTypeByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("contacttype.info.error id=%s: %v", id, err)
		return nil, err
	}

	return item, nil
}
