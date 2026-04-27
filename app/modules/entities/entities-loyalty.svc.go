package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.PointTransactionEntity = (*Service)(nil)

func (s *Service) ListPointTransactions(ctx context.Context) ([]*ent.PointTransaction, error) {
	var items []*ent.PointTransaction
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetPointTransactionByID(ctx context.Context, id uuid.UUID) (*ent.PointTransaction, error) {
	item := &ent.PointTransaction{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreatePointTransaction(ctx context.Context, item *ent.PointTransaction) (*ent.PointTransaction, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdatePointTransactionByID(ctx context.Context, id uuid.UUID, item *ent.PointTransaction) (*ent.PointTransaction, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetPointTransactionByID(ctx, id)
}

func (s *Service) DeletePointTransactionByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.PointTransaction)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
