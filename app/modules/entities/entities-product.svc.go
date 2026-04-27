package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductEntity = (*Service)(nil)

func (s *Service) ListProducts(ctx context.Context) ([]*ent.Product, error) {
	var items []*ent.Product
	err := s.db.NewSelect().
		Model(&items).
		Where("deleted_at IS NULL").
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetProductByID(ctx context.Context, id uuid.UUID) (*ent.Product, error) {
	item := &ent.Product{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	return item, err
}

func (s *Service) CreateProduct(ctx context.Context, p *ent.Product) (*ent.Product, error) {
	_, err := s.db.NewInsert().
		Model(p).
		Returning("*").
		Exec(ctx)
	return p, err
}

func (s *Service) UpdateProductByID(ctx context.Context, id uuid.UUID, p *ent.Product) (*ent.Product, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.Product)(nil)).
		Set("category_id = ?", p.CategoryID).
		Set("brand_id = ?", p.BrandID).
		Set("merchant_id = ?", p.MerchantID).
		Set("sku = ?", p.SKU).
		Set("name_th = ?", p.NameTh).
		Set("name_en = ?", p.NameEn).
		Set("short_description_th = ?", p.ShortDescriptionTh).
		Set("full_description_th = ?", p.FullDescriptionTh).
		Set("price = ?", p.Price).
		Set("discount_price = ?", p.DiscountPrice).
		Set("is_on_sale = ?", p.IsOnSale).
		Set("slug = ?", p.Slug).
		Set("meta_title = ?", p.MetaTitle).
		Set("meta_description = ?", p.MetaDescription).
		Set("status = ?", p.Status).
		Set("is_active = ?", p.IsActive).
		Set("is_featured = ?", p.IsFeatured).
		Set("weight = ?", p.Weight).
		Set("width = ?", p.Width).
		Set("length = ?", p.Length).
		Set("height = ?", p.Height).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetProductByID(ctx, id)
}

func (s *Service) DeleteProductByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewUpdate().
		Model((*ent.Product)(nil)).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
