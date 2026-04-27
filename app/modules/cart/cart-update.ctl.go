package cart

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
	MemberID   *string  `json:"member_id"`
	TotalItems *int     `json:"total_items"`
	TotalPrice *float64 `json:"total_price"`
}

func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.CartInvalidID, nil)
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
			base.BadRequest(ctx, i18n.CartNotFound, nil)
			return
		}
		log.Errf("cart.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.CartUpdateFailed, nil)
		return
	}

	input := *current
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.MemberID = &v
	}
	if body.TotalItems != nil {
		input.TotalItems = *body.TotalItems
	}
	if body.TotalPrice != nil {
		input.TotalPrice = *body.TotalPrice
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("cart.update.error: %v", err)
		base.InternalServerError(ctx, i18n.CartUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
