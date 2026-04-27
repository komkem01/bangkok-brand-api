package member

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Member, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "member.Info")
	defer span.End()

	member, err := s.db.GetMemberByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("member.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("member.info.ok id=%s", id)
	return member, nil
}
