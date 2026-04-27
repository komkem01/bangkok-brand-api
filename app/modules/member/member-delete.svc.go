package member

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "member.Delete")
	defer span.End()

	if err := s.db.DeleteMemberByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("member.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("member.delete.ok id=%s", id)
	return nil
}
