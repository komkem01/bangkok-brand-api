package contact

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.MemberContact) (*ent.MemberContact, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contact.Update")
	defer span.End()

	updated, err := s.db.UpdateContactByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("contact.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("contact.update.ok id=%s", id)
	return updated, nil
}
