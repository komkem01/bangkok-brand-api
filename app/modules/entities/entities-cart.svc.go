package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.CartEntity = (*Service)(nil)

func (s *Service) ListCarts(ctx context.Context) ([]*ent.Cart, error) {
	var items []*ent.Cart
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetCartByID(ctx context.Context, id uuid.UUID) (*ent.Cart, error) {
	item := &ent.Cart{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateCart(ctx context.Context, p *ent.Cart) (*ent.Cart, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateCartByID(ctx context.Context, id uuid.UUID, p *ent.Cart) (*ent.Cart, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.Cart)(nil)).
		Set("member_id = ?", p.MemberID).
		Set("total_items = ?", p.TotalItems).
		Set("total_price = ?", p.TotalPrice).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetCartByID(ctx, id)
}

func (s *Service) DeleteCartByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Cart)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
