package variant

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ProductVariant) (*ent.ProductVariant, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "variant.Create")
	defer span.End()
	created, err := s.db.CreateVariant(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("variant.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
