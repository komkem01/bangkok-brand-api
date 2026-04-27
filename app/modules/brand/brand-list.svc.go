package brand

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Brand, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "brand.List")
	defer span.End()

	items, err := s.db.ListBrands(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("brand.list.error: %v", err)
		return nil, err
	}

	log.Infof("brand.list.ok count=%d", len(items))
	return items, nil
}
