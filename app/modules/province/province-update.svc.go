package province

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	Name     string
	IsActive bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.Province, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "province.Update")
	defer span.End()

	province, err := s.db.UpdateProvinceByID(ctx, id, input.Name, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("province.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("province.update.ok id=%s", id)
	return province, nil
}
