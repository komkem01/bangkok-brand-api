package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ChatMessageEntity = (*Service)(nil)

func (s *Service) ListChatMessages(ctx context.Context) ([]*ent.ChatMessage, error) {
	var items []*ent.ChatMessage
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetChatMessageByID(ctx context.Context, id uuid.UUID) (*ent.ChatMessage, error) {
	item := &ent.ChatMessage{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateChatMessage(ctx context.Context, item *ent.ChatMessage) (*ent.ChatMessage, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateChatMessageByID(ctx context.Context, id uuid.UUID, item *ent.ChatMessage) (*ent.ChatMessage, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetChatMessageByID(ctx, id)
}

func (s *Service) DeleteChatMessageByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ChatMessage)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
