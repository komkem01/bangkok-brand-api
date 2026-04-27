package memberdevice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.MemberDevice) (*ent.MemberDevice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberdevice.Update")
	defer span.End()

	updated, err := s.db.UpdateMemberDeviceByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("memberdevice.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("memberdevice.update.ok id=%s", id)
	return updated, nil
}
