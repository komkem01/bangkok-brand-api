package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductAttributeEntity = (*Service)(nil)

func (s *Service) ListProductAttributes(ctx context.Context) ([]*ent.ProductAttribute, error) {
	var items []*ent.ProductAttribute
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductAttributeByID(ctx context.Context, id uuid.UUID) (*ent.ProductAttribute, error) {
	item := &ent.ProductAttribute{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductAttribute(ctx context.Context, p *ent.ProductAttribute) (*ent.ProductAttribute, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateProductAttributeByID(ctx context.Context, id uuid.UUID, p *ent.ProductAttribute) (*ent.ProductAttribute, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.ProductAttribute)(nil)).
		Set("name_th = ?", p.NameTh).
		Set("name_en = ?", p.NameEn).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductAttributeByID(ctx, id)
}

func (s *Service) DeleteProductAttributeByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductAttribute)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
