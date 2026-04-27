package contact

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.MemberContact, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contact.List")
	defer span.End()

	contacts, err := s.db.ListContacts(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("contact.list.error: %v", err)
		return nil, err
	}

	log.Infof("contact.list.ok count=%d", len(contacts))
	return contacts, nil
}
