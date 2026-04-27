package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.FlashSaleEventEntity = (*Service)(nil)

func (s *Service) ListFlashSaleEvents(ctx context.Context) ([]*ent.FlashSaleEvent, error) {
	var items []*ent.FlashSaleEvent
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetFlashSaleEventByID(ctx context.Context, id uuid.UUID) (*ent.FlashSaleEvent, error) {
	item := &ent.FlashSaleEvent{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateFlashSaleEvent(ctx context.Context, item *ent.FlashSaleEvent) (*ent.FlashSaleEvent, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateFlashSaleEventByID(ctx context.Context, id uuid.UUID, item *ent.FlashSaleEvent) (*ent.FlashSaleEvent, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetFlashSaleEventByID(ctx, id)
}

func (s *Service) DeleteFlashSaleEventByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.FlashSaleEvent)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
