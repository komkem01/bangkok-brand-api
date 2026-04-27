package variant

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ProductVariant, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "variant.Info")
	defer span.End()
	item, err := s.db.GetVariantByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("variant.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
