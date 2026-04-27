package brand

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Brand) (*ent.Brand, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "brand.Create")
	defer span.End()

	created, err := s.db.CreateBrand(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("brand.create.error: %v", err)
		return nil, err
	}

	log.Infof("brand.create.ok id=%s", created.ID)
	return created, nil
}
