package wishlist

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Wishlist) (*ent.Wishlist, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "wishlist.Create")
	defer span.End()
	created, err := s.db.CreateWishlist(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("wishlist.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
