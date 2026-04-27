package productattributevalue

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
	ProductID       *string  `json:"product_id"`
	AttributeID     *string  `json:"attribute_id"`
	ValueTh         *string  `json:"value_th"`
	ValueEn         *string  `json:"value_en"`
	AdditionalPrice *float64 `json:"additional_price"`
}

// Update godoc
// PATCH /product-attribute-values/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ProductAttributeValueInvalidID, nil)
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
			base.BadRequest(ctx, i18n.ProductAttributeValueNotFound, nil)
			return
		}
		log.Errf("product-attribute-value.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeValueUpdateFailed, nil)
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
	if body.AttributeID != nil {
		v, err := uuid.Parse(*body.AttributeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.AttributeID = &v
	}
	if body.ValueTh != nil {
		input.ValueTh = body.ValueTh
	}
	if body.ValueEn != nil {
		input.ValueEn = body.ValueEn
	}
	if body.AdditionalPrice != nil {
		input.AdditionalPrice = *body.AdditionalPrice
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("product-attribute-value.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeValueUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
