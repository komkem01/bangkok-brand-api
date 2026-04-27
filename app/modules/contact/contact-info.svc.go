package contact

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.MemberContact, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contact.Info")
	defer span.End()

	item, err := s.db.GetContactByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("contact.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("contact.info.ok id=%s", id)
	return item, nil
}
