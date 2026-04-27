package product

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	CategoryID         *string  `json:"category_id"`
	BrandID            *string  `json:"brand_id"`
	MerchantID         *string  `json:"merchant_id"`
	SKU                *string  `json:"sku"`
	NameTh             *string  `json:"name_th"`
	NameEn             *string  `json:"name_en"`
	ShortDescriptionTh *string  `json:"short_description_th"`
	FullDescriptionTh  *string  `json:"full_description_th"`
	Price              *float64 `json:"price"`
	DiscountPrice      *float64 `json:"discount_price"`
	IsOnSale           *bool    `json:"is_on_sale"`
	Slug               *string  `json:"slug"`
	MetaTitle          *string  `json:"meta_title"`
	MetaDescription    *string  `json:"meta_description"`
	Status             *string  `json:"status"`
	IsActive           *bool    `json:"is_active"`
	IsFeatured         *bool    `json:"is_featured"`
	Weight             *float64 `json:"weight"`
	Width              *float64 `json:"width"`
	Length             *float64 `json:"length"`
	Height             *float64 `json:"height"`
}

// Create godoc
// POST /products
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.Product{
		SKU:                body.SKU,
		NameTh:             body.NameTh,
		NameEn:             body.NameEn,
		ShortDescriptionTh: body.ShortDescriptionTh,
		FullDescriptionTh:  body.FullDescriptionTh,
		Price:              body.Price,
		DiscountPrice:      body.DiscountPrice,
		Slug:               body.Slug,
		MetaTitle:          body.MetaTitle,
		MetaDescription:    body.MetaDescription,
		Status:             body.Status,
		Weight:             body.Weight,
		Width:              body.Width,
		Length:             body.Length,
		Height:             body.Height,
		IsOnSale:           false,
		IsActive:           true,
		IsFeatured:         false,
	}
	if body.CategoryID != nil {
		v, err := uuid.Parse(*body.CategoryID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.CategoryID = &v
	}
	if body.BrandID != nil {
		v, err := uuid.Parse(*body.BrandID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.BrandID = &v
	}
	if body.MerchantID != nil {
		v, err := uuid.Parse(*body.MerchantID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.MerchantID = &v
	}
	if body.IsOnSale != nil {
		item.IsOnSale = *body.IsOnSale
	}
	if body.IsActive != nil {
		item.IsActive = *body.IsActive
	}
	if body.IsFeatured != nil {
		item.IsFeatured = *body.IsFeatured
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("product.create.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
