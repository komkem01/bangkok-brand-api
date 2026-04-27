package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.GenderEntity = (*Service)(nil)

// ListGenders returns all active (non-deleted) genders ordered by created_at.
func (s *Service) ListGenders(ctx context.Context) ([]*ent.Gender, error) {
	var genders []*ent.Gender
	err := s.db.NewSelect().
		Model(&genders).
		Where("deleted_at IS NULL").
		OrderExpr("created_at ASC").
		Scan(ctx)
	return genders, err
}

// GetGenderByID returns a single gender by its primary key.
func (s *Service) GetGenderByID(ctx context.Context, id uuid.UUID) (*ent.Gender, error) {
	gender := &ent.Gender{}
	err := s.db.NewSelect().
		Model(gender).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	return gender, err
}

// UpdateGenderByID updates name_th, name_en, is_active of a gender record.
func (s *Service) UpdateGenderByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.Gender, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Gender)(nil)).
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
	return s.GetGenderByID(ctx, id)
}

// DeleteGenderByID soft-deletes a gender record.
func (s *Service) DeleteGenderByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewUpdate().
		Model((*ent.Gender)(nil)).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
