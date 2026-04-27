package product

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBodyRequest struct {
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

// Update godoc
// PATCH /products/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ProductInvalidID, nil)
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	id := uuid.MustParse(uri.ID)
	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.ProductNotFound, nil)
			return
		}
		log.Errf("product.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductUpdateFailed, nil)
		return
	}

	input := *current
	if body.CategoryID != nil {
		v, err := uuid.Parse(*body.CategoryID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.CategoryID = &v
	}
	if body.BrandID != nil {
		v, err := uuid.Parse(*body.BrandID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.BrandID = &v
	}
	if body.MerchantID != nil {
		v, err := uuid.Parse(*body.MerchantID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.MerchantID = &v
	}
	if body.SKU != nil {
		input.SKU = body.SKU
	}
	if body.NameTh != nil {
		input.NameTh = body.NameTh
	}
	if body.NameEn != nil {
		input.NameEn = body.NameEn
	}
	if body.ShortDescriptionTh != nil {
		input.ShortDescriptionTh = body.ShortDescriptionTh
	}
	if body.FullDescriptionTh != nil {
		input.FullDescriptionTh = body.FullDescriptionTh
	}
	if body.Price != nil {
		input.Price = body.Price
	}
	if body.DiscountPrice != nil {
		input.DiscountPrice = body.DiscountPrice
	}
	if body.IsOnSale != nil {
		input.IsOnSale = *body.IsOnSale
	}
	if body.Slug != nil {
		input.Slug = body.Slug
	}
	if body.MetaTitle != nil {
		input.MetaTitle = body.MetaTitle
	}
	if body.MetaDescription != nil {
		input.MetaDescription = body.MetaDescription
	}
	if body.Status != nil {
		input.Status = body.Status
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}
	if body.IsFeatured != nil {
		input.IsFeatured = *body.IsFeatured
	}
	if body.Weight != nil {
		input.Weight = body.Weight
	}
	if body.Width != nil {
		input.Width = body.Width
	}
	if body.Length != nil {
		input.Length = body.Length
	}
	if body.Height != nil {
		input.Height = body.Height
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("product.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
