package product

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID                 string   `json:"id"`
	CategoryID         *string  `json:"category_id,omitempty"`
	BrandID            *string  `json:"brand_id,omitempty"`
	MerchantID         *string  `json:"merchant_id,omitempty"`
	SKU                *string  `json:"sku,omitempty"`
	NameTh             *string  `json:"name_th,omitempty"`
	NameEn             *string  `json:"name_en,omitempty"`
	ShortDescriptionTh *string  `json:"short_description_th,omitempty"`
	FullDescriptionTh  *string  `json:"full_description_th,omitempty"`
	Price              *float64 `json:"price,omitempty"`
	DiscountPrice      *float64 `json:"discount_price,omitempty"`
	IsOnSale           bool     `json:"is_on_sale"`
	Slug               *string  `json:"slug,omitempty"`
	MetaTitle          *string  `json:"meta_title,omitempty"`
	MetaDescription    *string  `json:"meta_description,omitempty"`
	Status             *string  `json:"status,omitempty"`
	IsActive           bool     `json:"is_active"`
	IsFeatured         bool     `json:"is_featured"`
	Weight             *float64 `json:"weight,omitempty"`
	Width              *float64 `json:"width,omitempty"`
	Length             *float64 `json:"length,omitempty"`
	Height             *float64 `json:"height,omitempty"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}

func toListResponse(p *ent.Product) ListResponse {
	return ListResponse{
		ID:                 p.ID.String(),
		CategoryID:         uuidToStringPtr(p.CategoryID),
		BrandID:            uuidToStringPtr(p.BrandID),
		MerchantID:         uuidToStringPtr(p.MerchantID),
		SKU:                p.SKU,
		NameTh:             p.NameTh,
		NameEn:             p.NameEn,
		ShortDescriptionTh: p.ShortDescriptionTh,
		FullDescriptionTh:  p.FullDescriptionTh,
		Price:              p.Price,
		DiscountPrice:      p.DiscountPrice,
		IsOnSale:           p.IsOnSale,
		Slug:               p.Slug,
		MetaTitle:          p.MetaTitle,
		MetaDescription:    p.MetaDescription,
		Status:             p.Status,
		IsActive:           p.IsActive,
		IsFeatured:         p.IsFeatured,
		Weight:             p.Weight,
		Width:              p.Width,
		Length:             p.Length,
		Height:             p.Height,
		CreatedAt:          p.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:          p.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func uuidToStringPtr(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	s := id.String()
	return &s
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("product.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
