package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShippingZoneAreaEntity = (*Service)(nil)

func (s *Service) ListShippingZoneAreas(ctx context.Context) ([]*ent.ShippingZoneArea, error) {
	var items []*ent.ShippingZoneArea
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShippingZoneAreaByID(ctx context.Context, id uuid.UUID) (*ent.ShippingZoneArea, error) {
	item := &ent.ShippingZoneArea{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShippingZoneArea(ctx context.Context, item *ent.ShippingZoneArea) (*ent.ShippingZoneArea, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShippingZoneAreaByID(ctx context.Context, id uuid.UUID, item *ent.ShippingZoneArea) (*ent.ShippingZoneArea, error) {
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShippingZoneAreaByID(ctx, id)
}

func (s *Service) DeleteShippingZoneAreaByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShippingZoneArea)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
