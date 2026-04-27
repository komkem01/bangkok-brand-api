package contact

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
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

func (s *Service) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberContact, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contact.ListByMemberID")
	defer span.End()

	contacts, err := s.db.ListContactsByMemberID(ctx, memberID)
	if err != nil {
		span.RecordError(err)
		log.Errf("contact.list_by_member.error member_id=%s: %v", memberID, err)
		return nil, err
	}

	log.Infof("contact.list_by_member.ok member_id=%s count=%d", memberID, len(contacts))
	return contacts, nil
}
