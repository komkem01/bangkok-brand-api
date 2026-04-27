package zipcode

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	SubDistrictID *uuid.UUID
	Name          string
	IsActive      bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.Zipcode, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "zipcode.Update")
	defer span.End()

	zipcode, err := s.db.UpdateZipcodeByID(ctx, id, input.SubDistrictID, input.Name, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("zipcode.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("zipcode.update.ok id=%s", id)
	return zipcode, nil
}
