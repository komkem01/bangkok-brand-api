package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ReturnRequestEntity = (*Service)(nil)

func (s *Service) ListReturnRequests(ctx context.Context) ([]*ent.ReturnRequest, error) {
	var items []*ent.ReturnRequest
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetReturnRequestByID(ctx context.Context, id uuid.UUID) (*ent.ReturnRequest, error) {
	item := &ent.ReturnRequest{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateReturnRequest(ctx context.Context, item *ent.ReturnRequest) (*ent.ReturnRequest, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateReturnRequestByID(ctx context.Context, id uuid.UUID, item *ent.ReturnRequest) (*ent.ReturnRequest, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetReturnRequestByID(ctx, id)
}

func (s *Service) DeleteReturnRequestByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ReturnRequest)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
