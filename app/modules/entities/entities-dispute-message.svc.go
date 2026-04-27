package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.DisputeMessageEntity = (*Service)(nil)

func (s *Service) ListDisputeMessages(ctx context.Context) ([]*ent.DisputeMessage, error) {
	var items []*ent.DisputeMessage
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetDisputeMessageByID(ctx context.Context, id uuid.UUID) (*ent.DisputeMessage, error) {
	item := &ent.DisputeMessage{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateDisputeMessage(ctx context.Context, item *ent.DisputeMessage) (*ent.DisputeMessage, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateDisputeMessageByID(ctx context.Context, id uuid.UUID, item *ent.DisputeMessage) (*ent.DisputeMessage, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetDisputeMessageByID(ctx, id)
}

func (s *Service) DeleteDisputeMessageByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.DisputeMessage)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
