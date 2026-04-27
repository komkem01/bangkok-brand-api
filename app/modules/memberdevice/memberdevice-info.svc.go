package memberdevice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.MemberDevice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberdevice.Info")
	defer span.End()

	item, err := s.db.GetMemberDeviceByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("memberdevice.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("memberdevice.info.ok id=%s", id)
	return item, nil
}
