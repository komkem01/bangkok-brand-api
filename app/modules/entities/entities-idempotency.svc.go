package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.IdempotencyKeyEntity = (*Service)(nil)

func (s *Service) ListIdempotencyKeys(ctx context.Context) ([]*ent.IdempotencyKey, error) {
	var items []*ent.IdempotencyKey
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetIdempotencyKeyByID(ctx context.Context, id uuid.UUID) (*ent.IdempotencyKey, error) {
	item := &ent.IdempotencyKey{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateIdempotencyKey(ctx context.Context, item *ent.IdempotencyKey) (*ent.IdempotencyKey, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateIdempotencyKeyByID(ctx context.Context, id uuid.UUID, item *ent.IdempotencyKey) (*ent.IdempotencyKey, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetIdempotencyKeyByID(ctx, id)
}

func (s *Service) DeleteIdempotencyKeyByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.IdempotencyKey)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
