package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ChatRoomEntity = (*Service)(nil)

func (s *Service) ListChatRooms(ctx context.Context) ([]*ent.ChatRoom, error) {
	var items []*ent.ChatRoom
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetChatRoomByID(ctx context.Context, id uuid.UUID) (*ent.ChatRoom, error) {
	item := &ent.ChatRoom{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateChatRoom(ctx context.Context, item *ent.ChatRoom) (*ent.ChatRoom, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateChatRoomByID(ctx context.Context, id uuid.UUID, item *ent.ChatRoom) (*ent.ChatRoom, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetChatRoomByID(ctx, id)
}

func (s *Service) DeleteChatRoomByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ChatRoom)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
