package district

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.District, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "district.List")
	defer span.End()

	districts, err := s.db.ListDistricts(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("district.list.error: %v", err)
		return nil, err
	}

	log.Infof("district.list.ok count=%d", len(districts))
	return districts, nil
}
