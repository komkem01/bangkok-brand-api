package productstock

import (
	"database/sql"
	"errors"
	"time"

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
	ProductID         *string `json:"product_id"`
	Quantity          *int    `json:"quantity"`
	LowStockThreshold *int    `json:"low_stock_threshold"`
	LastRestockedAt   *string `json:"last_restocked_at"`
}

// Update godoc
// PATCH /product-stocks/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ProductStockInvalidID, nil)
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
			base.BadRequest(ctx, i18n.ProductStockNotFound, nil)
			return
		}
		log.Errf("product-stock.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductStockUpdateFailed, nil)
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
	if body.Quantity != nil {
		input.Quantity = *body.Quantity
	}
	if body.LowStockThreshold != nil {
		input.LowStockThreshold = *body.LowStockThreshold
	}
	if body.LastRestockedAt != nil {
		t, err := time.Parse(time.RFC3339, *body.LastRestockedAt)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.LastRestockedAt = &t
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("product-stock.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductStockUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
