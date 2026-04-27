package gender

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "gender.Delete")
	defer span.End()

	if err := s.db.DeleteGenderByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("gender.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("gender.delete.ok id=%s", id)
	return nil
}
