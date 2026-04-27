package province

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Province, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "province.Info")
	defer span.End()

	province, err := s.db.GetProvinceByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("province.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("province.info.ok id=%s", id)
	return province, nil
}
