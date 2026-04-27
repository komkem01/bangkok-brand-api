package address

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) List(ctx context.Context) ([]*ent.MemberAddress, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.List")
	defer span.End()

	items, err := s.db.ListAddresses(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("address.list.error: %v", err)
		return nil, err
	}

	log.Infof("address.list.ok count=%d", len(items))
	return items, nil
}

func (s *Service) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberAddress, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.ListByMemberID")
	defer span.End()

	items, err := s.db.ListAddressesByMemberID(ctx, memberID)
	if err != nil {
		span.RecordError(err)
		log.Errf("address.list_by_member.error member_id=%s: %v", memberID, err)
		return nil, err
	}

	log.Infof("address.list_by_member.ok member_id=%s count=%d", memberID, len(items))
	return items, nil
}
