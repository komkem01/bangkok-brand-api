package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.WebhookEventEntity = (*Service)(nil)

func (s *Service) ListWebhookEvents(ctx context.Context) ([]*ent.WebhookEvent, error) {
	var items []*ent.WebhookEvent
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetWebhookEventByID(ctx context.Context, id uuid.UUID) (*ent.WebhookEvent, error) {
	item := &ent.WebhookEvent{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateWebhookEvent(ctx context.Context, item *ent.WebhookEvent) (*ent.WebhookEvent, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateWebhookEventByID(ctx context.Context, id uuid.UUID, item *ent.WebhookEvent) (*ent.WebhookEvent, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetWebhookEventByID(ctx, id)
}

func (s *Service) DeleteWebhookEventByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.WebhookEvent)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
