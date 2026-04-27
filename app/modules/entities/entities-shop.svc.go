package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShopEntity = (*Service)(nil)

func (s *Service) ListShops(ctx context.Context) ([]*ent.Shop, error) {
	var items []*ent.Shop
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetShopByID(ctx context.Context, id uuid.UUID) (*ent.Shop, error) {
	item := &ent.Shop{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateShop(ctx context.Context, item *ent.Shop) (*ent.Shop, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateShopByID(ctx context.Context, id uuid.UUID, item *ent.Shop) (*ent.Shop, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShopByID(ctx, id)
}

func (s *Service) DeleteShopByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Shop)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
