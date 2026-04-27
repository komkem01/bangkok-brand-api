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

type ListResponse struct {
	ID                string  `json:"id"`
	ProductID         *string `json:"product_id,omitempty"`
	Quantity          int     `json:"quantity"`
	LowStockThreshold int     `json:"low_stock_threshold"`
	LastRestockedAt   *string `json:"last_restocked_at,omitempty"`
	UpdatedAt         string  `json:"updated_at"`
}

func toListResponse(p *ent.ProductStock) ListResponse {
	return ListResponse{
		ID:                p.ID.String(),
		ProductID:         uuidToStringPtr(p.ProductID),
		Quantity:          p.Quantity,
		LowStockThreshold: p.LowStockThreshold,
		LastRestockedAt:   timeToStringPtr(p.LastRestockedAt),
		UpdatedAt:         p.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func uuidToStringPtr(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	s := id.String()
	return &s
}

func timeToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format(utils.RFC3339Milli)
	return &s
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("product-stock.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductStockListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
