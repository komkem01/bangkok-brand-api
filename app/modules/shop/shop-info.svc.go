package shop

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Shop, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shop.Info")
	defer span.End()
	item, err := s.db.GetShopByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("shop.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
