package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.InvoiceEntity = (*Service)(nil)

func (s *Service) ListInvoices(ctx context.Context) ([]*ent.Invoice, error) {
	var items []*ent.Invoice
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error) {
	item := &ent.Invoice{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateInvoice(ctx context.Context, item *ent.Invoice) (*ent.Invoice, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateInvoiceByID(ctx context.Context, id uuid.UUID, item *ent.Invoice) (*ent.Invoice, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetInvoiceByID(ctx, id)
}

func (s *Service) DeleteInvoiceByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Invoice)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
