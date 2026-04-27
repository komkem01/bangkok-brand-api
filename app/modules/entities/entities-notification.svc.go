package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.NotificationEntity = (*Service)(nil)

func (s *Service) ListNotifications(ctx context.Context) ([]*ent.Notification, error) {
	var items []*ent.Notification
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetNotificationByID(ctx context.Context, id uuid.UUID) (*ent.Notification, error) {
	item := &ent.Notification{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateNotification(ctx context.Context, item *ent.Notification) (*ent.Notification, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateNotificationByID(ctx context.Context, id uuid.UUID, item *ent.Notification) (*ent.Notification, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetNotificationByID(ctx, id)
}

func (s *Service) DeleteNotificationByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Notification)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
