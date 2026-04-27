package productimage

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	ProductID *string `json:"product_id"`
	StorageID *string `json:"storage_id"`
	IsMain    *bool   `json:"is_main"`
	SortOrder *int    `json:"sort_order"`
}

// Create godoc
// POST /product-images
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.ProductImage{IsMain: false, SortOrder: 0}
	if body.ProductID != nil {
		v, err := uuid.Parse(*body.ProductID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ProductID = &v
	}
	if body.StorageID != nil {
		v, err := uuid.Parse(*body.StorageID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.StorageID = &v
	}
	if body.IsMain != nil {
		item.IsMain = *body.IsMain
	}
	if body.SortOrder != nil {
		item.SortOrder = *body.SortOrder
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("product-image.create.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductImageCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
