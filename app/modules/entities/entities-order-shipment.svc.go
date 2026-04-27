package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.OrderShipmentEntity = (*Service)(nil)

func (s *Service) ListOrderShipments(ctx context.Context) ([]*ent.OrderShipment, error) {
	var items []*ent.OrderShipment
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetOrderShipmentByID(ctx context.Context, id uuid.UUID) (*ent.OrderShipment, error) {
	item := &ent.OrderShipment{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateOrderShipment(ctx context.Context, item *ent.OrderShipment) (*ent.OrderShipment, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateOrderShipmentByID(ctx context.Context, id uuid.UUID, item *ent.OrderShipment) (*ent.OrderShipment, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetOrderShipmentByID(ctx, id)
}

func (s *Service) DeleteOrderShipmentByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.OrderShipment)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
