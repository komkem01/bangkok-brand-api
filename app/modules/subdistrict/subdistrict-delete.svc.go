package subdistrict

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "subdistrict.Delete")
	defer span.End()

	if err := s.db.DeleteSubdistrictByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("subdistrict.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("subdistrict.delete.ok id=%s", id)
	return nil
}
