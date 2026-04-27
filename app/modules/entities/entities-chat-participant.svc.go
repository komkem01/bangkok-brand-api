package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ChatParticipantEntity = (*Service)(nil)

func (s *Service) ListChatParticipants(ctx context.Context) ([]*ent.ChatParticipant, error) {
	var items []*ent.ChatParticipant
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("id DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetChatParticipantByID(ctx context.Context, id uuid.UUID) (*ent.ChatParticipant, error) {
	item := &ent.ChatParticipant{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateChatParticipant(ctx context.Context, item *ent.ChatParticipant) (*ent.ChatParticipant, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateChatParticipantByID(ctx context.Context, id uuid.UUID, item *ent.ChatParticipant) (*ent.ChatParticipant, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetChatParticipantByID(ctx, id)
}

func (s *Service) DeleteChatParticipantByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ChatParticipant)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
