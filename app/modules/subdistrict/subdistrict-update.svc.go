package subdistrict

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	DistrictID *uuid.UUID
	Name       string
	IsActive   bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.Subdistrict, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "subdistrict.Update")
	defer span.End()

	subdistrict, err := s.db.UpdateSubdistrictByID(ctx, id, input.DistrictID, input.Name, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("subdistrict.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("subdistrict.update.ok id=%s", id)
	return subdistrict, nil
}
