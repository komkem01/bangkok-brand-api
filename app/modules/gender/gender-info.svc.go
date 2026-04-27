package gender

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Gender, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "gender.Info")
	defer span.End()

	gender, err := s.db.GetGenderByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("gender.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("gender.info.ok id=%s", id)
	return gender, nil
}
