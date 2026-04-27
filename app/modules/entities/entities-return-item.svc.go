package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ReturnItemEntity = (*Service)(nil)

func (s *Service) ListReturnItems(ctx context.Context) ([]*ent.ReturnItem, error) {
	var items []*ent.ReturnItem
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetReturnItemByID(ctx context.Context, id uuid.UUID) (*ent.ReturnItem, error) {
	item := &ent.ReturnItem{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateReturnItem(ctx context.Context, item *ent.ReturnItem) (*ent.ReturnItem, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateReturnItemByID(ctx context.Context, id uuid.UUID, item *ent.ReturnItem) (*ent.ReturnItem, error) {
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetReturnItemByID(ctx, id)
}

func (s *Service) DeleteReturnItemByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ReturnItem)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
