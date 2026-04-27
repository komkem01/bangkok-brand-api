package member

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Member, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "member.List")
	defer span.End()

	members, err := s.db.ListMembers(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("member.list.error: %v", err)
		return nil, err
	}

	log.Infof("member.list.ok count=%d", len(members))
	return members, nil
}
