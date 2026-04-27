package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.SettlementItemEntity = (*Service)(nil)

func (s *Service) ListSettlementItems(ctx context.Context) ([]*ent.SettlementItem, error) {
	var items []*ent.SettlementItem
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetSettlementItemByID(ctx context.Context, id uuid.UUID) (*ent.SettlementItem, error) {
	item := &ent.SettlementItem{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateSettlementItem(ctx context.Context, item *ent.SettlementItem) (*ent.SettlementItem, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateSettlementItemByID(ctx context.Context, id uuid.UUID, item *ent.SettlementItem) (*ent.SettlementItem, error) {
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetSettlementItemByID(ctx, id)
}

func (s *Service) DeleteSettlementItemByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.SettlementItem)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
