package cart

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID         string  `json:"id"`
	MemberID   *string `json:"member_id,omitempty"`
	TotalItems int     `json:"total_items"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

func toListResponse(p *ent.Cart) ListResponse {
	return ListResponse{
		ID:         p.ID.String(),
		MemberID:   uuidToStringPtr(p.MemberID),
		TotalItems: p.TotalItems,
		TotalPrice: p.TotalPrice,
		CreatedAt:  p.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:  p.UpdatedAt.Format(utils.RFC3339Milli),
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
		log.Errf("cart.list.error: %v", err)
		base.InternalServerError(ctx, i18n.CartListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
