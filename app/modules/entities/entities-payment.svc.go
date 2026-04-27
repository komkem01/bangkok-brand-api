package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.PaymentEntity = (*Service)(nil)

func (s *Service) ListPayments(ctx context.Context) ([]*ent.Payment, error) {
	var items []*ent.Payment
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetPaymentByID(ctx context.Context, id uuid.UUID) (*ent.Payment, error) {
	item := &ent.Payment{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreatePayment(ctx context.Context, p *ent.Payment) (*ent.Payment, error) {
	_, err := s.db.NewInsert().Model(p).Returning("*").Exec(ctx)
	return p, err
}

func (s *Service) UpdatePaymentByID(ctx context.Context, id uuid.UUID, p *ent.Payment) (*ent.Payment, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.Payment)(nil)).
		Set("order_id = ?", p.OrderID).
		Set("payment_no = ?", p.PaymentNo).
		Set("method = ?", p.Method).
		Set("amount = ?", p.Amount).
		Set("status = ?", p.Status).
		Set("evidence_storage_id = ?", p.EvidenceStorageID).
		Set("transfer_date_time = ?", p.TransferDateTime).
		Set("from_bank_id = ?", p.FromBankID).
		Set("transaction_ref = ?", p.TransactionRef).
		Set("paid_at = ?", p.PaidAt).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetPaymentByID(ctx, id)
}

func (s *Service) DeletePaymentByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.Payment)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
