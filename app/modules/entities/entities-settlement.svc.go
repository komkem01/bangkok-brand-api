package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.SettlementEntity = (*Service)(nil)

func (s *Service) ListSettlements(ctx context.Context) ([]*ent.SettlementBatch, error) {
	var items []*ent.SettlementBatch
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetSettlementByID(ctx context.Context, id uuid.UUID) (*ent.SettlementBatch, error) {
	item := &ent.SettlementBatch{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateSettlement(ctx context.Context, item *ent.SettlementBatch) (*ent.SettlementBatch, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateSettlementByID(ctx context.Context, id uuid.UUID, item *ent.SettlementBatch) (*ent.SettlementBatch, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetSettlementByID(ctx, id)
}

func (s *Service) DeleteSettlementByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.SettlementBatch)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
