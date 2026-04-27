package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.FlashSaleItemEntity = (*Service)(nil)

func (s *Service) ListFlashSaleItems(ctx context.Context) ([]*ent.FlashSaleItem, error) {
	var items []*ent.FlashSaleItem
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetFlashSaleItemByID(ctx context.Context, id uuid.UUID) (*ent.FlashSaleItem, error) {
	item := &ent.FlashSaleItem{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateFlashSaleItem(ctx context.Context, item *ent.FlashSaleItem) (*ent.FlashSaleItem, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateFlashSaleItemByID(ctx context.Context, id uuid.UUID, item *ent.FlashSaleItem) (*ent.FlashSaleItem, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetFlashSaleItemByID(ctx, id)
}

func (s *Service) DeleteFlashSaleItemByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.FlashSaleItem)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
