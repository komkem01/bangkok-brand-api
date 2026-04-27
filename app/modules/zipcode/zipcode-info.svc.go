package zipcode

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Zipcode, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "zipcode.Info")
	defer span.End()

	zipcode, err := s.db.GetZipcodeByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("zipcode.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("zipcode.info.ok id=%s", id)
	return zipcode, nil
}
