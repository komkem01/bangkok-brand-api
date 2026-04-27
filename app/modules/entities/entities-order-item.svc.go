package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.OrderItemEntity = (*Service)(nil)

func (s *Service) ListOrderItems(ctx context.Context) ([]*ent.OrderItem, error) {
	var items []*ent.OrderItem
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetOrderItemByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	item := &ent.OrderItem{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateOrderItem(ctx context.Context, p *ent.OrderItem) (*ent.OrderItem, error) {
	_, err := s.db.NewInsert().Model(p).Returning("*").Exec(ctx)
	return p, err
}

func (s *Service) UpdateOrderItemByID(ctx context.Context, id uuid.UUID, p *ent.OrderItem) (*ent.OrderItem, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.OrderItem)(nil)).
		Set("order_id = ?", p.OrderID).
		Set("product_id = ?", p.ProductID).
		Set("product_name_snapshot = ?", p.ProductNameSnapshot).
		Set("sku_snapshot = ?", p.SKUSnapshot).
		Set("quantity = ?", p.Quantity).
		Set("unit_price = ?", p.UnitPrice).
		Set("selected_attributes = ?", p.SelectedAttributes).
		Set("subtotal_price = ?", p.SubtotalPrice).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetOrderItemByID(ctx, id)
}

func (s *Service) DeleteOrderItemByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.OrderItem)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
