package gender

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	NameTh   string
	NameEn   string
	IsActive bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.Gender, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "gender.Update")
	defer span.End()

	gender, err := s.db.UpdateGenderByID(ctx, id, input.NameTh, input.NameEn, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("gender.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("gender.update.ok id=%s", id)
	return gender, nil
}
