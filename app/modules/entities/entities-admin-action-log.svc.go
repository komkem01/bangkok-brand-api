package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.AdminActionLogEntity = (*Service)(nil)

func (s *Service) ListAdminActionLogs(ctx context.Context) ([]*ent.AdminActionLog, error) {
	var items []*ent.AdminActionLog
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetAdminActionLogByID(ctx context.Context, id uuid.UUID) (*ent.AdminActionLog, error) {
	item := &ent.AdminActionLog{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateAdminActionLog(ctx context.Context, item *ent.AdminActionLog) (*ent.AdminActionLog, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateAdminActionLogByID(ctx context.Context, id uuid.UUID, item *ent.AdminActionLog) (*ent.AdminActionLog, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetAdminActionLogByID(ctx, id)
}

func (s *Service) DeleteAdminActionLogByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.AdminActionLog)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
