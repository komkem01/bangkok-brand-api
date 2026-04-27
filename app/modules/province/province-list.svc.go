package province

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Province, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "province.List")
	defer span.End()

	provinces, err := s.db.ListProvinces(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("province.list.error: %v", err)
		return nil, err
	}

	log.Infof("province.list.ok count=%d", len(provinces))
	return provinces, nil
}
