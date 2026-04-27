package shop

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "shop.Delete")
	defer span.End()
	err := s.db.DeleteShopByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("shop.delete.error: %v", err)
	}
	return err
}
