package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.LogisticsProviderEntity = (*Service)(nil)

func (s *Service) ListLogisticsProviders(ctx context.Context) ([]*ent.LogisticsProvider, error) {
	var items []*ent.LogisticsProvider
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetLogisticsProviderByID(ctx context.Context, id uuid.UUID) (*ent.LogisticsProvider, error) {
	item := &ent.LogisticsProvider{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateLogisticsProvider(ctx context.Context, item *ent.LogisticsProvider) (*ent.LogisticsProvider, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateLogisticsProviderByID(ctx context.Context, id uuid.UUID, item *ent.LogisticsProvider) (*ent.LogisticsProvider, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetLogisticsProviderByID(ctx, id)
}

func (s *Service) DeleteLogisticsProviderByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.LogisticsProvider)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
