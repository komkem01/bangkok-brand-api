package productattribute

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
	NameTh *string `json:"name_th"`
	NameEn *string `json:"name_en"`
}

// Update godoc
// PATCH /product-attributes/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ProductAttributeInvalidID, nil)
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
			base.BadRequest(ctx, i18n.ProductAttributeNotFound, nil)
			return
		}
		log.Errf("product-attribute.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeUpdateFailed, nil)
		return
	}

	input := *current
	if body.NameTh != nil {
		input.NameTh = body.NameTh
	}
	if body.NameEn != nil {
		input.NameEn = body.NameEn
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("product-attribute.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
