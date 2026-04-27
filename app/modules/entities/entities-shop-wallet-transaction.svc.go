package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShopWalletTransactionEntity = (*Service)(nil)

func (s *Service) ListShopWalletTransactions(ctx context.Context) ([]*ent.ShopWalletTransaction, error) {
	var items []*ent.ShopWalletTransaction
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShopWalletTransactionByID(ctx context.Context, id uuid.UUID) (*ent.ShopWalletTransaction, error) {
	item := &ent.ShopWalletTransaction{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShopWalletTransaction(ctx context.Context, item *ent.ShopWalletTransaction) (*ent.ShopWalletTransaction, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShopWalletTransactionByID(ctx context.Context, id uuid.UUID, item *ent.ShopWalletTransaction) (*ent.ShopWalletTransaction, error) {
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShopWalletTransactionByID(ctx, id)
}

func (s *Service) DeleteShopWalletTransactionByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShopWalletTransaction)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
