package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShopSettingEntity = (*Service)(nil)

func (s *Service) ListShopSettings(ctx context.Context) ([]*ent.ShopSetting, error) {
	var items []*ent.ShopSetting
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShopSettingByID(ctx context.Context, id uuid.UUID) (*ent.ShopSetting, error) {
	item := &ent.ShopSetting{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShopSetting(ctx context.Context, item *ent.ShopSetting) (*ent.ShopSetting, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShopSettingByID(ctx context.Context, id uuid.UUID, item *ent.ShopSetting) (*ent.ShopSetting, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShopSettingByID(ctx, id)
}

func (s *Service) DeleteShopSettingByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShopSetting)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
