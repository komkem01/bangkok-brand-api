package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.RewardEntity = (*Service)(nil)

func (s *Service) ListRewards(ctx context.Context) ([]*ent.Reward, error) {
	var items []*ent.Reward
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetRewardByID(ctx context.Context, id uuid.UUID) (*ent.Reward, error) {
	item := &ent.Reward{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateReward(ctx context.Context, item *ent.Reward) (*ent.Reward, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateRewardByID(ctx context.Context, id uuid.UUID, item *ent.Reward) (*ent.Reward, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetRewardByID(ctx, id)
}

func (s *Service) DeleteRewardByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.Reward)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
