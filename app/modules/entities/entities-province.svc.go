package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProvinceEntity = (*Service)(nil)

// ListProvinces returns all active provinces ordered by name ascending.
func (s *Service) ListProvinces(ctx context.Context) ([]*ent.Province, error) {
	var provinces []*ent.Province
	err := s.db.NewSelect().
		Model(&provinces).
		Where("is_active = true").
		OrderExpr("name ASC").
		OrderExpr("id ASC").
		Scan(ctx)
	return provinces, err
}

// GetProvinceByID returns a single province by its primary key.
func (s *Service) GetProvinceByID(ctx context.Context, id uuid.UUID) (*ent.Province, error) {
	province := &ent.Province{}
	err := s.db.NewSelect().
		Model(province).
		Where("id = ?", id).
		Scan(ctx)
	return province, err
}

// UpdateProvinceByID updates name and is_active of a province record.
func (s *Service) UpdateProvinceByID(ctx context.Context, id uuid.UUID, name string, isActive bool) (*ent.Province, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Province)(nil)).
		Set("name = ?", name).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProvinceByID(ctx, id)
}

// DeleteProvinceByID hard-deletes a province record.
func (s *Service) DeleteProvinceByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Province)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
