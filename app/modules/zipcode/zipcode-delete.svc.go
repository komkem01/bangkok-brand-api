package zipcode

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "zipcode.Delete")
	defer span.End()

	if err := s.db.DeleteZipcodeByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("zipcode.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("zipcode.delete.ok id=%s", id)
	return nil
}
