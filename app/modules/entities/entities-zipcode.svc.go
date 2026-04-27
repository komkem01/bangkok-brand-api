package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ZipcodeEntity = (*Service)(nil)

// ListZipcodes returns all active zipcodes ordered by created_at.
func (s *Service) ListZipcodes(ctx context.Context) ([]*ent.Zipcode, error) {
	var zipcodes []*ent.Zipcode
	err := s.db.NewSelect().
		Model(&zipcodes).
		Where("is_active = true").
		OrderExpr("created_at ASC").
		Scan(ctx)
	return zipcodes, err
}

// GetZipcodeByID returns a single zipcode by its primary key.
func (s *Service) GetZipcodeByID(ctx context.Context, id uuid.UUID) (*ent.Zipcode, error) {
	zipcode := &ent.Zipcode{}
	err := s.db.NewSelect().
		Model(zipcode).
		Where("id = ?", id).
		Scan(ctx)
	return zipcode, err
}

// UpdateZipcodeByID updates sub_district_id, name, and is_active of a zipcode record.
func (s *Service) UpdateZipcodeByID(ctx context.Context, id uuid.UUID, subDistrictID *uuid.UUID, name string, isActive bool) (*ent.Zipcode, error) {
	now := time.Now()
	q := s.db.NewUpdate().
		Model((*ent.Zipcode)(nil)).
		Set("name = ?", name).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id)
	if subDistrictID != nil {
		q = q.Set("sub_district_id = ?", *subDistrictID)
	}
	if _, err := q.Exec(ctx); err != nil {
		return nil, err
	}
	return s.GetZipcodeByID(ctx, id)
}

// DeleteZipcodeByID hard-deletes a zipcode record.
func (s *Service) DeleteZipcodeByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Zipcode)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
