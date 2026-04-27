package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.PointSettingEntity = (*Service)(nil)

func (s *Service) ListPointSettings(ctx context.Context) ([]*ent.PointSetting, error) {
	var items []*ent.PointSetting
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("updated_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetPointSettingByID(ctx context.Context, id uuid.UUID) (*ent.PointSetting, error) {
	item := &ent.PointSetting{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreatePointSetting(ctx context.Context, item *ent.PointSetting) (*ent.PointSetting, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdatePointSettingByID(ctx context.Context, id uuid.UUID, item *ent.PointSetting) (*ent.PointSetting, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetPointSettingByID(ctx, id)
}

func (s *Service) DeletePointSettingByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.PointSetting)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
