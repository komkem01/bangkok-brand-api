package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.SubdistrictEntity = (*Service)(nil)

// ListSubdistricts returns all active sub-districts ordered by district and sub-district name ascending.
func (s *Service) ListSubdistricts(ctx context.Context) ([]*ent.Subdistrict, error) {
	var subdistricts []*ent.Subdistrict
	err := s.db.NewSelect().
		Model(&subdistricts).
		Where("is_active = true").
		OrderExpr("district_id ASC").
		OrderExpr("name ASC").
		OrderExpr("id ASC").
		Scan(ctx)
	return subdistricts, err
}

// GetSubdistrictByID returns a single sub-district by its primary key.
func (s *Service) GetSubdistrictByID(ctx context.Context, id uuid.UUID) (*ent.Subdistrict, error) {
	subdistrict := &ent.Subdistrict{}
	err := s.db.NewSelect().
		Model(subdistrict).
		Where("id = ?", id).
		Scan(ctx)
	return subdistrict, err
}

// UpdateSubdistrictByID updates district_id, name, and is_active of a sub-district record.
func (s *Service) UpdateSubdistrictByID(ctx context.Context, id uuid.UUID, districtID *uuid.UUID, name string, isActive bool) (*ent.Subdistrict, error) {
	now := time.Now()
	q := s.db.NewUpdate().
		Model((*ent.Subdistrict)(nil)).
		Set("name = ?", name).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id)
	if districtID != nil {
		q = q.Set("district_id = ?", *districtID)
	}
	if _, err := q.Exec(ctx); err != nil {
		return nil, err
	}
	return s.GetSubdistrictByID(ctx, id)
}

// DeleteSubdistrictByID hard-deletes a sub-district record.
func (s *Service) DeleteSubdistrictByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Subdistrict)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
