package province

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "province.Delete")
	defer span.End()

	if err := s.db.DeleteProvinceByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("province.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("province.delete.ok id=%s", id)
	return nil
}
