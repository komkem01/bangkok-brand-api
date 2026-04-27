package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.OrderStatusHistoryEntity = (*Service)(nil)

func (s *Service) ListOrderStatusHistories(ctx context.Context) ([]*ent.OrderStatusHistory, error) {
	var items []*ent.OrderStatusHistory
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetOrderStatusHistoryByID(ctx context.Context, id uuid.UUID) (*ent.OrderStatusHistory, error) {
	item := &ent.OrderStatusHistory{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateOrderStatusHistory(ctx context.Context, p *ent.OrderStatusHistory) (*ent.OrderStatusHistory, error) {
	_, err := s.db.NewInsert().Model(p).Returning("*").Exec(ctx)
	return p, err
}

func (s *Service) UpdateOrderStatusHistoryByID(ctx context.Context, id uuid.UUID, p *ent.OrderStatusHistory) (*ent.OrderStatusHistory, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.OrderStatusHistory)(nil)).
		Set("order_id = ?", p.OrderID).
		Set("status = ?", p.Status).
		Set("remark = ?", p.Remark).
		Set("changed_by_id = ?", p.ChangedByID).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetOrderStatusHistoryByID(ctx, id)
}

func (s *Service) DeleteOrderStatusHistoryByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.OrderStatusHistory)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
