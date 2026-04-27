package district

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.District, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "district.Info")
	defer span.End()

	district, err := s.db.GetDistrictByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("district.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("district.info.ok id=%s", id)
	return district, nil
}
