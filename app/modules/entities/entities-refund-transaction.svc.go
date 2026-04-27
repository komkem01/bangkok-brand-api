package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.RefundTransactionEntity = (*Service)(nil)

func (s *Service) ListRefundTransactions(ctx context.Context) ([]*ent.RefundTransaction, error) {
	var items []*ent.RefundTransaction
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetRefundTransactionByID(ctx context.Context, id uuid.UUID) (*ent.RefundTransaction, error) {
	item := &ent.RefundTransaction{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateRefundTransaction(ctx context.Context, item *ent.RefundTransaction) (*ent.RefundTransaction, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateRefundTransactionByID(ctx context.Context, id uuid.UUID, item *ent.RefundTransaction) (*ent.RefundTransaction, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetRefundTransactionByID(ctx, id)
}

func (s *Service) DeleteRefundTransactionByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.RefundTransaction)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
