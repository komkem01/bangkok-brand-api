package cart

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	MemberID   *string  `json:"member_id"`
	TotalItems *int     `json:"total_items"`
	TotalPrice *float64 `json:"total_price"`
}

func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.Cart{TotalItems: 0, TotalPrice: 0}
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.MemberID = &v
	}
	if body.TotalItems != nil {
		item.TotalItems = *body.TotalItems
	}
	if body.TotalPrice != nil {
		item.TotalPrice = *body.TotalPrice
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("cart.create.error: %v", err)
		base.InternalServerError(ctx, i18n.CartCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
