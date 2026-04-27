package variant

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ProductVariant, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "variant.List")
	defer span.End()
	items, err := s.db.ListVariants(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("variant.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
