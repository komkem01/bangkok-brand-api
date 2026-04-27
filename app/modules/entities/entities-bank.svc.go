package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.BankEntity = (*Service)(nil)

// ListBanks returns all active banks ordered by name_th ascending.
func (s *Service) ListBanks(ctx context.Context) ([]*ent.Bank, error) {
	var banks []*ent.Bank
	err := s.db.NewSelect().
		Model(&banks).
		Where("is_active = true").
		OrderExpr("name_th ASC").
		OrderExpr("id ASC").
		Scan(ctx)
	return banks, err
}

// GetBankByID returns a single bank by its primary key.
func (s *Service) GetBankByID(ctx context.Context, id uuid.UUID) (*ent.Bank, error) {
	bank := &ent.Bank{}
	err := s.db.NewSelect().
		Model(bank).
		Where("id = ?", id).
		Scan(ctx)
	return bank, err
}

func (s *Service) CreateBank(ctx context.Context, b *ent.Bank) (*ent.Bank, error) {
	_, err := s.db.NewInsert().
		Model(b).
		Returning("*").
		Exec(ctx)
	return b, err
}

// UpdateBankByID updates fields of a bank record.
func (s *Service) UpdateBankByID(ctx context.Context, id uuid.UUID, nameTh, nameEn, code string, isActive bool) (*ent.Bank, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Bank)(nil)).
		Set("name_th = ?", nameTh).
		Set("name_en = ?", nameEn).
		Set("code = ?", code).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetBankByID(ctx, id)
}

// DeleteBankByID hard-deletes a bank record.
func (s *Service) DeleteBankByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Bank)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
