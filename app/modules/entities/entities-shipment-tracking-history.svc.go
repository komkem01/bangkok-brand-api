package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShipmentTrackingHistoryEntity = (*Service)(nil)

func (s *Service) ListShipmentTrackingHistories(ctx context.Context) ([]*ent.ShipmentTrackingHistory, error) {
	var items []*ent.ShipmentTrackingHistory
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShipmentTrackingHistoryByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentTrackingHistory, error) {
	item := &ent.ShipmentTrackingHistory{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShipmentTrackingHistory(ctx context.Context, item *ent.ShipmentTrackingHistory) (*ent.ShipmentTrackingHistory, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShipmentTrackingHistoryByID(ctx context.Context, id uuid.UUID, item *ent.ShipmentTrackingHistory) (*ent.ShipmentTrackingHistory, error) {
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShipmentTrackingHistoryByID(ctx, id)
}

func (s *Service) DeleteShipmentTrackingHistoryByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShipmentTrackingHistory)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
