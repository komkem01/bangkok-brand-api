package zipcode

import (
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (c *Controller) Delete(ctx *gin.Context) {
	var req DeleteRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.ZipcodeInvalidID, nil)
		return
	}

	id := uuid.MustParse(req.ID)
	if err := c.svc.Delete(ctx.Request.Context(), id); err != nil {
		log.Errf("zipcode.delete.error: %v", err)
		base.InternalServerError(ctx, i18n.ZipcodeDeleteFailed, nil)
		return
	}

	base.Success(ctx, nil)
}
