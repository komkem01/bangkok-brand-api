package productstock

import (
	"time"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	ProductID         *string `json:"product_id"`
	Quantity          *int    `json:"quantity"`
	LowStockThreshold *int    `json:"low_stock_threshold"`
	LastRestockedAt   *string `json:"last_restocked_at"`
}

// Create godoc
// POST /product-stocks
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.ProductStock{Quantity: 0, LowStockThreshold: 5}
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
	if body.LowStockThreshold != nil {
		item.LowStockThreshold = *body.LowStockThreshold
	}
	if body.LastRestockedAt != nil {
		t, err := time.Parse(time.RFC3339, *body.LastRestockedAt)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.LastRestockedAt = &t
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("product-stock.create.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductStockCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
