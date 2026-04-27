package cartitem

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
	CartID                  *string        `json:"cart_id"`
	ProductID               *string        `json:"product_id"`
	Quantity                *int           `json:"quantity"`
	SelectedAttributeValues map[string]any `json:"selected_attribute_values"`
	UnitPrice               *float64       `json:"unit_price"`
	SubtotalPrice           *float64       `json:"subtotal_price"`
}

func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.CartItemInvalidID, nil)
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
			base.BadRequest(ctx, i18n.CartItemNotFound, nil)
			return
		}
		log.Errf("cart-item.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.CartItemUpdateFailed, nil)
		return
	}

	input := *current
	if body.CartID != nil {
		v, err := uuid.Parse(*body.CartID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.CartID = &v
	}
	if body.ProductID != nil {
		v, err := uuid.Parse(*body.ProductID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ProductID = &v
	}
	if body.Quantity != nil {
		input.Quantity = *body.Quantity
	}
	if body.SelectedAttributeValues != nil {
		input.SelectedAttributeValues = body.SelectedAttributeValues
	}
	if body.UnitPrice != nil {
		input.UnitPrice = body.UnitPrice
	}
	if body.SubtotalPrice != nil {
		input.SubtotalPrice = body.SubtotalPrice
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("cart-item.update.error: %v", err)
		base.InternalServerError(ctx, i18n.CartItemUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
