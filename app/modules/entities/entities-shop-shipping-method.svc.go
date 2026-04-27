package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShopShippingMethodEntity = (*Service)(nil)

func (s *Service) ListShopShippingMethods(ctx context.Context) ([]*ent.ShopShippingMethod, error) {
	var items []*ent.ShopShippingMethod
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShopShippingMethodByID(ctx context.Context, id uuid.UUID) (*ent.ShopShippingMethod, error) {
	item := &ent.ShopShippingMethod{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShopShippingMethod(ctx context.Context, item *ent.ShopShippingMethod) (*ent.ShopShippingMethod, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShopShippingMethodByID(ctx context.Context, id uuid.UUID, item *ent.ShopShippingMethod) (*ent.ShopShippingMethod, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShopShippingMethodByID(ctx, id)
}

func (s *Service) DeleteShopShippingMethodByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShopShippingMethod)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
