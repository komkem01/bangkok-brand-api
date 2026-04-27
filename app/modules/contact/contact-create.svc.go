package contact

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.MemberContact) (*ent.MemberContact, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contact.Create")
	defer span.End()

	created, err := s.db.CreateContact(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("contact.create.error: %v", err)
		return nil, err
	}

	log.Infof("contact.create.ok id=%s", created.ID)
	return created, nil
}
