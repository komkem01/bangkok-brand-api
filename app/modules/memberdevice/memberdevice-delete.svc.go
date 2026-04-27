package memberdevice

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberdevice.Delete")
	defer span.End()

	if err := s.db.DeleteMemberDeviceByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("memberdevice.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("memberdevice.delete.ok id=%s", id)
	return nil
}
