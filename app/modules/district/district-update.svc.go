package district

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	ProvinceID *uuid.UUID
	Name       string
	IsActive   bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.District, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "district.Update")
	defer span.End()

	district, err := s.db.UpdateDistrictByID(ctx, id, input.ProvinceID, input.Name, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("district.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("district.update.ok id=%s", id)
	return district, nil
}
