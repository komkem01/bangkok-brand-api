package member

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, member *ent.Member) (*ent.Member, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "member.Update")
	defer span.End()

	updated, err := s.db.UpdateMemberByID(ctx, id, member)
	if err != nil {
		span.RecordError(err)
		log.Errf("member.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("member.update.ok id=%s", id)
	return updated, nil
}
