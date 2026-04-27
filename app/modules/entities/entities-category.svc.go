package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.CategoryEntity = (*Service)(nil)

func (s *Service) ListCategories(ctx context.Context) ([]*ent.Category, error) {
	var items []*ent.Category
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("sort_order ASC").
		OrderExpr("created_at ASC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetCategoryByID(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	item := &ent.Category{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateCategory(ctx context.Context, c *ent.Category) (*ent.Category, error) {
	_, err := s.db.NewInsert().
		Model(c).
		Returning("*").
		Exec(ctx)
	return c, err
}

func (s *Service) UpdateCategoryByID(ctx context.Context, id uuid.UUID, c *ent.Category) (*ent.Category, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Category)(nil)).
		Set("parent_id = ?", c.ParentID).
		Set("name_th = ?", c.NameTh).
		Set("name_en = ?", c.NameEn).
		Set("description = ?", c.Description).
		Set("image_id = ?", c.ImageID).
		Set("slug = ?", c.Slug).
		Set("is_active = ?", c.IsActive).
		Set("sort_order = ?", c.SortOrder).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetCategoryByID(ctx, id)
}

func (s *Service) DeleteCategoryByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.Category)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
