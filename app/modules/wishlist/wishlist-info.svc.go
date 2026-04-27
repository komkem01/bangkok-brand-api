package wishlist

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Wishlist, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "wishlist.Info")
	defer span.End()
	item, err := s.db.GetWishlistByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("wishlist.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
