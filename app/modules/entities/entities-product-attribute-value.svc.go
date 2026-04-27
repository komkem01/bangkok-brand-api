package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductAttributeValueEntity = (*Service)(nil)

func (s *Service) ListProductAttributeValues(ctx context.Context) ([]*ent.ProductAttributeValue, error) {
	var items []*ent.ProductAttributeValue
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("id DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductAttributeValueByID(ctx context.Context, id uuid.UUID) (*ent.ProductAttributeValue, error) {
	item := &ent.ProductAttributeValue{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProductAttributeValue(ctx context.Context, p *ent.ProductAttributeValue) (*ent.ProductAttributeValue, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateProductAttributeValueByID(ctx context.Context, id uuid.UUID, p *ent.ProductAttributeValue) (*ent.ProductAttributeValue, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.ProductAttributeValue)(nil)).
		Set("product_id = ?", p.ProductID).
		Set("attribute_id = ?", p.AttributeID).
		Set("value_th = ?", p.ValueTh).
		Set("value_en = ?", p.ValueEn).
		Set("additional_price = ?", p.AdditionalPrice).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductAttributeValueByID(ctx, id)
}

func (s *Service) DeleteProductAttributeValueByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ProductAttributeValue)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
