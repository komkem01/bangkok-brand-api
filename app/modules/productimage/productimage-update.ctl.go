package productimage

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
	ProductID *string `json:"product_id"`
	StorageID *string `json:"storage_id"`
	IsMain    *bool   `json:"is_main"`
	SortOrder *int    `json:"sort_order"`
}

// Update godoc
// PATCH /product-images/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ProductImageInvalidID, nil)
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
			base.BadRequest(ctx, i18n.ProductImageNotFound, nil)
			return
		}
		log.Errf("product-image.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductImageUpdateFailed, nil)
		return
	}

	input := *current
	if body.ProductID != nil {
		v, err := uuid.Parse(*body.ProductID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ProductID = &v
	}
	if body.StorageID != nil {
		v, err := uuid.Parse(*body.StorageID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.StorageID = &v
	}
	if body.IsMain != nil {
		input.IsMain = *body.IsMain
	}
	if body.SortOrder != nil {
		input.SortOrder = *body.SortOrder
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("product-image.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductImageUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
