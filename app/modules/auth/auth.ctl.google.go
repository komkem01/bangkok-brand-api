package auth

import (
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GoogleLogin(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)
	defer span.End()
	base.NotImplemented(ctx, i18n.BadRequest, nil)
}

func (c *Controller) GoogleCallback(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)
	defer span.End()
	base.NotImplemented(ctx, i18n.BadRequest, nil)
}
