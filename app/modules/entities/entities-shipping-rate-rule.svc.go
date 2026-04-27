package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShippingRateRuleEntity = (*Service)(nil)

func (s *Service) ListShippingRateRules(ctx context.Context) ([]*ent.ShippingRateRule, error) {
	var items []*ent.ShippingRateRule
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShippingRateRuleByID(ctx context.Context, id uuid.UUID) (*ent.ShippingRateRule, error) {
	item := &ent.ShippingRateRule{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShippingRateRule(ctx context.Context, item *ent.ShippingRateRule) (*ent.ShippingRateRule, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShippingRateRuleByID(ctx context.Context, id uuid.UUID, item *ent.ShippingRateRule) (*ent.ShippingRateRule, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShippingRateRuleByID(ctx, id)
}

func (s *Service) DeleteShippingRateRuleByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShippingRateRule)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
