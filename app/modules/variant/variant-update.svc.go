package variant

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.ProductVariant) (*ent.ProductVariant, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "variant.Update")
	defer span.End()
	updated, err := s.db.UpdateVariantByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("variant.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
