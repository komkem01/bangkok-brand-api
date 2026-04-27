package productattributevalue

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	ProductID       *string  `json:"product_id"`
	AttributeID     *string  `json:"attribute_id"`
	ValueTh         *string  `json:"value_th"`
	ValueEn         *string  `json:"value_en"`
	AdditionalPrice *float64 `json:"additional_price"`
}

// Create godoc
// POST /product-attribute-values
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.ProductAttributeValue{
		ValueTh:         body.ValueTh,
		ValueEn:         body.ValueEn,
		AdditionalPrice: 0,
	}
	if body.ProductID != nil {
		v, err := uuid.Parse(*body.ProductID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ProductID = &v
	}
	if body.AttributeID != nil {
		v, err := uuid.Parse(*body.AttributeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.AttributeID = &v
	}
	if body.AdditionalPrice != nil {
		item.AdditionalPrice = *body.AdditionalPrice
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("product-attribute-value.create.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeValueCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
