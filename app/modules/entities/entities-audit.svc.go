package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.AuditLogEntity = (*Service)(nil)

func (s *Service) ListAuditLogs(ctx context.Context) ([]*ent.AuditLog, error) {
	var items []*ent.AuditLog
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetAuditLogByID(ctx context.Context, id uuid.UUID) (*ent.AuditLog, error) {
	item := &ent.AuditLog{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateAuditLog(ctx context.Context, item *ent.AuditLog) (*ent.AuditLog, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateAuditLogByID(ctx context.Context, id uuid.UUID, item *ent.AuditLog) (*ent.AuditLog, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetAuditLogByID(ctx, id)
}

func (s *Service) DeleteAuditLogByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.AuditLog)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
