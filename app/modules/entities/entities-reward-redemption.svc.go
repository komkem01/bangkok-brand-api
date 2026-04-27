package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.RewardRedemptionEntity = (*Service)(nil)

func (s *Service) ListRewardRedemptions(ctx context.Context) ([]*ent.RewardRedemption, error) {
	var items []*ent.RewardRedemption
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetRewardRedemptionByID(ctx context.Context, id uuid.UUID) (*ent.RewardRedemption, error) {
	item := &ent.RewardRedemption{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateRewardRedemption(ctx context.Context, item *ent.RewardRedemption) (*ent.RewardRedemption, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateRewardRedemptionByID(ctx context.Context, id uuid.UUID, item *ent.RewardRedemption) (*ent.RewardRedemption, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetRewardRedemptionByID(ctx, id)
}

func (s *Service) DeleteRewardRedemptionByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.RewardRedemption)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
