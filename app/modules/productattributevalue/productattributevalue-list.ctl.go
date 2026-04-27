package productattributevalue

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID              string  `json:"id"`
	ProductID       *string `json:"product_id,omitempty"`
	AttributeID     *string `json:"attribute_id,omitempty"`
	ValueTh         *string `json:"value_th,omitempty"`
	ValueEn         *string `json:"value_en,omitempty"`
	AdditionalPrice float64 `json:"additional_price"`
}

func toListResponse(p *ent.ProductAttributeValue) ListResponse {
	return ListResponse{
		ID:              p.ID.String(),
		ProductID:       uuidToStringPtr(p.ProductID),
		AttributeID:     uuidToStringPtr(p.AttributeID),
		ValueTh:         p.ValueTh,
		ValueEn:         p.ValueEn,
		AdditionalPrice: p.AdditionalPrice,
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
		log.Errf("product-attribute-value.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeValueListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
