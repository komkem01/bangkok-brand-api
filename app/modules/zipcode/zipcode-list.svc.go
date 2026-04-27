package zipcode

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Zipcode, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "zipcode.List")
	defer span.End()

	zipcodes, err := s.db.ListZipcodes(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("zipcode.list.error: %v", err)
		return nil, err
	}

	log.Infof("zipcode.list.ok count=%d", len(zipcodes))
	return zipcodes, nil
}
