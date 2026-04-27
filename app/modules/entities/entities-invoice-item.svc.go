package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.InvoiceItemEntity = (*Service)(nil)

func (s *Service) ListInvoiceItems(ctx context.Context) ([]*ent.InvoiceItem, error) {
	var items []*ent.InvoiceItem
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetInvoiceItemByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceItem, error) {
	item := &ent.InvoiceItem{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateInvoiceItem(ctx context.Context, item *ent.InvoiceItem) (*ent.InvoiceItem, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateInvoiceItemByID(ctx context.Context, id uuid.UUID, item *ent.InvoiceItem) (*ent.InvoiceItem, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetInvoiceItemByID(ctx, id)
}

func (s *Service) DeleteInvoiceItemByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.InvoiceItem)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
