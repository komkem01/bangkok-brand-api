package auth

import (
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Logout(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := c.svc.Logout(ctx.Request.Context()); err != nil {
		log.Errf("auth.logout.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}

	base.Success(ctx, nil)
}
