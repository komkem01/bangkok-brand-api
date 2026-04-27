package bank

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InfoRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (c *Controller) Info(ctx *gin.Context) {
	var req InfoRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BankInvalidID, nil)
		return
	}

	id := uuid.MustParse(req.ID)
	bank, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.BankNotFound, nil)
			return
		}
		log.Errf("bank.info.error: %v", err)
		base.InternalServerError(ctx, i18n.BankInfoFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(bank))
}
