package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.KYCVerificationEntity = (*Service)(nil)

func (s *Service) ListKYCVerifications(ctx context.Context) ([]*ent.KYCVerification, error) {
	var items []*ent.KYCVerification
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetKYCVerificationByID(ctx context.Context, id uuid.UUID) (*ent.KYCVerification, error) {
	item := &ent.KYCVerification{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateKYCVerification(ctx context.Context, item *ent.KYCVerification) (*ent.KYCVerification, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateKYCVerificationByID(ctx context.Context, id uuid.UUID, item *ent.KYCVerification) (*ent.KYCVerification, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetKYCVerificationByID(ctx, id)
}

func (s *Service) DeleteKYCVerificationByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.KYCVerification)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
