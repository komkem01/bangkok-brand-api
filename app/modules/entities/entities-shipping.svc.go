package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShippingZoneEntity = (*Service)(nil)

func (s *Service) ListShippingZones(ctx context.Context) ([]*ent.ShippingZone, error) {
	var items []*ent.ShippingZone
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetShippingZoneByID(ctx context.Context, id uuid.UUID) (*ent.ShippingZone, error) {
	item := &ent.ShippingZone{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateShippingZone(ctx context.Context, item *ent.ShippingZone) (*ent.ShippingZone, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateShippingZoneByID(ctx context.Context, id uuid.UUID, item *ent.ShippingZone) (*ent.ShippingZone, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShippingZoneByID(ctx, id)
}

func (s *Service) DeleteShippingZoneByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ShippingZone)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
