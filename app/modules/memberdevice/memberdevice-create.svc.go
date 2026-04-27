package memberdevice

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.MemberDevice) (*ent.MemberDevice, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberdevice.Create")
	defer span.End()

	created, err := s.db.CreateMemberDevice(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("memberdevice.create.error: %v", err)
		return nil, err
	}

	log.Infof("memberdevice.create.ok id=%s", created.ID)
	return created, nil
}
