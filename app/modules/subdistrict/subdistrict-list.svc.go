package subdistrict

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Subdistrict, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "subdistrict.List")
	defer span.End()

	subdistricts, err := s.db.ListSubdistricts(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("subdistrict.list.error: %v", err)
		return nil, err
	}

	log.Infof("subdistrict.list.ok count=%d", len(subdistricts))
	return subdistricts, nil
}
