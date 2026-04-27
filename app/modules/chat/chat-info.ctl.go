package chat

import (
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) Info(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	id := uuid.MustParse(req.ID)
	item, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		log.Errf("chat.info.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, item)
}
