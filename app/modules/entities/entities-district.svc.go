package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.DistrictEntity = (*Service)(nil)

// ListDistricts returns all active districts ordered by created_at.
func (s *Service) ListDistricts(ctx context.Context) ([]*ent.District, error) {
	var districts []*ent.District
	err := s.db.NewSelect().
		Model(&districts).
		Where("is_active = true").
		OrderExpr("created_at ASC").
		Scan(ctx)
	return districts, err
}

// GetDistrictByID returns a single district by its primary key.
func (s *Service) GetDistrictByID(ctx context.Context, id uuid.UUID) (*ent.District, error) {
	district := &ent.District{}
	err := s.db.NewSelect().
		Model(district).
		Where("id = ?", id).
		Scan(ctx)
	return district, err
}

// UpdateDistrictByID updates province_id, name, and is_active of a district record.
func (s *Service) UpdateDistrictByID(ctx context.Context, id uuid.UUID, provinceID *uuid.UUID, name string, isActive bool) (*ent.District, error) {
	now := time.Now()
	q := s.db.NewUpdate().
		Model((*ent.District)(nil)).
		Set("name = ?", name).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id)
	if provinceID != nil {
		q = q.Set("province_id = ?", *provinceID)
	}
	if _, err := q.Exec(ctx); err != nil {
		return nil, err
	}
	return s.GetDistrictByID(ctx, id)
}

// DeleteDistrictByID hard-deletes a district record.
func (s *Service) DeleteDistrictByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.District)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
