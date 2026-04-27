package cartitem

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID                      string         `json:"id"`
	CartID                  *string        `json:"cart_id,omitempty"`
	ProductID               *string        `json:"product_id,omitempty"`
	Quantity                int            `json:"quantity"`
	SelectedAttributeValues map[string]any `json:"selected_attribute_values,omitempty"`
	UnitPrice               *float64       `json:"unit_price,omitempty"`
	SubtotalPrice           *float64       `json:"subtotal_price,omitempty"`
	CreatedAt               string         `json:"created_at"`
	UpdatedAt               string         `json:"updated_at"`
}

func toListResponse(p *ent.CartItem) ListResponse {
	return ListResponse{
		ID:                      p.ID.String(),
		CartID:                  uuidToStringPtr(p.CartID),
		ProductID:               uuidToStringPtr(p.ProductID),
		Quantity:                p.Quantity,
		SelectedAttributeValues: p.SelectedAttributeValues,
		UnitPrice:               p.UnitPrice,
		SubtotalPrice:           p.SubtotalPrice,
		CreatedAt:               p.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:               p.UpdatedAt.Format(utils.RFC3339Milli),
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
		log.Errf("cart-item.list.error: %v", err)
		base.InternalServerError(ctx, i18n.CartItemListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
