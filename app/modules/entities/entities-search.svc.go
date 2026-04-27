package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.SearchHistoryEntity = (*Service)(nil)

func (s *Service) ListSearchHistories(ctx context.Context) ([]*ent.SearchHistory, error) {
	var items []*ent.SearchHistory
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetSearchHistoryByID(ctx context.Context, id uuid.UUID) (*ent.SearchHistory, error) {
	item := &ent.SearchHistory{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateSearchHistory(ctx context.Context, item *ent.SearchHistory) (*ent.SearchHistory, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateSearchHistoryByID(ctx context.Context, id uuid.UUID, item *ent.SearchHistory) (*ent.SearchHistory, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetSearchHistoryByID(ctx, id)
}

func (s *Service) DeleteSearchHistoryByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.SearchHistory)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
