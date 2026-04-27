package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.WishlistEntity = (*Service)(nil)

func (s *Service) ListWishlists(ctx context.Context) ([]*ent.Wishlist, error) {
	var items []*ent.Wishlist
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetWishlistByID(ctx context.Context, id uuid.UUID) (*ent.Wishlist, error) {
	item := &ent.Wishlist{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateWishlist(ctx context.Context, item *ent.Wishlist) (*ent.Wishlist, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateWishlistByID(ctx context.Context, id uuid.UUID, item *ent.Wishlist) (*ent.Wishlist, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetWishlistByID(ctx, id)
}

func (s *Service) DeleteWishlistByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Wishlist)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
