package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.CouponEntity = (*Service)(nil)

func (s *Service) ListCoupons(ctx context.Context) ([]*ent.Coupon, error) {
	var items []*ent.Coupon
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetCouponByID(ctx context.Context, id uuid.UUID) (*ent.Coupon, error) {
	item := &ent.Coupon{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateCoupon(ctx context.Context, item *ent.Coupon) (*ent.Coupon, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateCouponByID(ctx context.Context, id uuid.UUID, item *ent.Coupon) (*ent.Coupon, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetCouponByID(ctx, id)
}

func (s *Service) DeleteCouponByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Coupon)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
