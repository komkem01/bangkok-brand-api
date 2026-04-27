package memberdevice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) List(ctx context.Context) ([]*ent.MemberDevice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberdevice.List")
	defer span.End()

	memberdevices, err := s.db.ListMemberDevices(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("memberdevice.list.error: %v", err)
		return nil, err
	}

	log.Infof("memberdevice.list.ok count=%d", len(memberdevices))
	return memberdevices, nil
}

func (s *Service) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberDevice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberdevice.ListByMemberID")
	defer span.End()

	memberdevices, err := s.db.ListMemberDevicesByMemberID(ctx, memberID)
	if err != nil {
		span.RecordError(err)
		log.Errf("memberdevice.list_by_member.error member_id=%s: %v", memberID, err)
		return nil, err
	}

	log.Infof("memberdevice.list_by_member.ok member_id=%s count=%d", memberID, len(memberdevices))
	return memberdevices, nil
}
