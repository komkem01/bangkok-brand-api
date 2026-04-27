package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.MemberBankAccountEntity = (*Service)(nil)

func (s *Service) ListMemberBankAccounts(ctx context.Context) ([]*ent.MemberBankAccount, error) {
	var items []*ent.MemberBankAccount
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetMemberBankAccountByID(ctx context.Context, id uuid.UUID) (*ent.MemberBankAccount, error) {
	item := &ent.MemberBankAccount{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateMemberBankAccount(ctx context.Context, a *ent.MemberBankAccount) (*ent.MemberBankAccount, error) {
	_, err := s.db.NewInsert().
		Model(a).
		Returning("*").
		Exec(ctx)
	return a, err
}

func (s *Service) UpdateMemberBankAccountByID(ctx context.Context, id uuid.UUID, a *ent.MemberBankAccount) (*ent.MemberBankAccount, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.MemberBankAccount)(nil)).
		Set("member_id = ?", a.MemberID).
		Set("bank_id = ?", a.BankID).
		Set("account_number = ?", a.AccountNumber).
		Set("account_name = ?", a.AccountName).
		Set("branch_name = ?", a.BranchName).
		Set("is_default = ?", a.IsDefault).
		Set("is_verified = ?", a.IsVerified).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetMemberBankAccountByID(ctx, id)
}

func (s *Service) DeleteMemberBankAccountByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.MemberBankAccount)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
