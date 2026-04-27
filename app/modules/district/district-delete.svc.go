package district

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "district.Delete")
	defer span.End()

	if err := s.db.DeleteDistrictByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("district.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("district.delete.ok id=%s", id)
	return nil
}
