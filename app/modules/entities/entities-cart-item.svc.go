package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.CartItemEntity = (*Service)(nil)

func (s *Service) ListCartItems(ctx context.Context) ([]*ent.CartItem, error) {
	var items []*ent.CartItem
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetCartItemByID(ctx context.Context, id uuid.UUID) (*ent.CartItem, error) {
	item := &ent.CartItem{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateCartItem(ctx context.Context, p *ent.CartItem) (*ent.CartItem, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateCartItemByID(ctx context.Context, id uuid.UUID, p *ent.CartItem) (*ent.CartItem, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.CartItem)(nil)).
		Set("cart_id = ?", p.CartID).
		Set("product_id = ?", p.ProductID).
		Set("quantity = ?", p.Quantity).
		Set("selected_attribute_values = ?", p.SelectedAttributeValues).
		Set("unit_price = ?", p.UnitPrice).
		Set("subtotal_price = ?", p.SubtotalPrice).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetCartItemByID(ctx, id)
}

func (s *Service) DeleteCartItemByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.CartItem)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
