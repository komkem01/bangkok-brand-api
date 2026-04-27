package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.CouponUsageEntity = (*Service)(nil)

func (s *Service) ListCouponUsages(ctx context.Context) ([]*ent.CouponUsage, error) {
	var items []*ent.CouponUsage
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetCouponUsageByID(ctx context.Context, id uuid.UUID) (*ent.CouponUsage, error) {
	item := &ent.CouponUsage{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateCouponUsage(ctx context.Context, item *ent.CouponUsage) (*ent.CouponUsage, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateCouponUsageByID(ctx context.Context, id uuid.UUID, item *ent.CouponUsage) (*ent.CouponUsage, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetCouponUsageByID(ctx, id)
}

func (s *Service) DeleteCouponUsageByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.CouponUsage)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
