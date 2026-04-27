package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.PrefixEntity = (*Service)(nil)

// ListPrefixes returns all active (non-deleted) prefixes ordered by created_at.
func (s *Service) ListPrefixes(ctx context.Context) ([]*ent.Prefix, error) {
	var prefixes []*ent.Prefix
	err := s.db.NewSelect().
		Model(&prefixes).
		Where("deleted_at IS NULL").
		OrderExpr("created_at ASC").
		Scan(ctx)
	return prefixes, err
}

// GetPrefixByID returns a single prefix by its primary key.
func (s *Service) GetPrefixByID(ctx context.Context, id uuid.UUID) (*ent.Prefix, error) {
	prefix := &ent.Prefix{}
	err := s.db.NewSelect().
		Model(prefix).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	return prefix, err
}

// UpdatePrefixByID updates fields of a prefix record.
func (s *Service) UpdatePrefixByID(ctx context.Context, id uuid.UUID, genderID *uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.Prefix, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Prefix)(nil)).
		Set("gender_id = ?", genderID).
		Set("name_th = ?", nameTh).
		Set("name_en = ?", nameEn).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetPrefixByID(ctx, id)
}

// DeletePrefixByID soft-deletes a prefix record.
func (s *Service) DeletePrefixByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewUpdate().
		Model((*ent.Prefix)(nil)).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
