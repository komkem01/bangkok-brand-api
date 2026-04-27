package cartitem

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	CartID                  *string        `json:"cart_id"`
	ProductID               *string        `json:"product_id"`
	Quantity                *int           `json:"quantity"`
	SelectedAttributeValues map[string]any `json:"selected_attribute_values"`
	UnitPrice               *float64       `json:"unit_price"`
	SubtotalPrice           *float64       `json:"subtotal_price"`
}

func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.CartItem{Quantity: 1, SelectedAttributeValues: body.SelectedAttributeValues, UnitPrice: body.UnitPrice, SubtotalPrice: body.SubtotalPrice}
	if body.CartID != nil {
		v, err := uuid.Parse(*body.CartID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.CartID = &v
	}
	if body.ProductID != nil {
		v, err := uuid.Parse(*body.ProductID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ProductID = &v
	}
	if body.Quantity != nil {
		item.Quantity = *body.Quantity
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("cart-item.create.error: %v", err)
		base.InternalServerError(ctx, i18n.CartItemCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
