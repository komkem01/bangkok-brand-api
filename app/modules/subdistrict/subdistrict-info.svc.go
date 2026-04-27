package subdistrict

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Subdistrict, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "subdistrict.Info")
	defer span.End()

	subdistrict, err := s.db.GetSubdistrictByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("subdistrict.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("subdistrict.info.ok id=%s", id)
	return subdistrict, nil
}
