package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.DisputeCaseEntity = (*Service)(nil)

func (s *Service) ListDisputeCases(ctx context.Context) ([]*ent.DisputeCase, error) {
	var items []*ent.DisputeCase
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetDisputeCaseByID(ctx context.Context, id uuid.UUID) (*ent.DisputeCase, error) {
	item := &ent.DisputeCase{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateDisputeCase(ctx context.Context, item *ent.DisputeCase) (*ent.DisputeCase, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateDisputeCaseByID(ctx context.Context, id uuid.UUID, item *ent.DisputeCase) (*ent.DisputeCase, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetDisputeCaseByID(ctx, id)
}

func (s *Service) DeleteDisputeCaseByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.DisputeCase)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
