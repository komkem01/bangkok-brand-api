package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.KYCStatusHistoryEntity = (*Service)(nil)

func (s *Service) ListKYCStatusHistories(ctx context.Context) ([]*ent.KYCStatusHistory, error) {
	var items []*ent.KYCStatusHistory
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetKYCStatusHistoryByID(ctx context.Context, id uuid.UUID) (*ent.KYCStatusHistory, error) {
	item := &ent.KYCStatusHistory{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateKYCStatusHistory(ctx context.Context, item *ent.KYCStatusHistory) (*ent.KYCStatusHistory, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateKYCStatusHistoryByID(ctx context.Context, id uuid.UUID, item *ent.KYCStatusHistory) (*ent.KYCStatusHistory, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetKYCStatusHistoryByID(ctx, id)
}

func (s *Service) DeleteKYCStatusHistoryByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.KYCStatusHistory)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
