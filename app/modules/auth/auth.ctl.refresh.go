package auth

import (
	"errors"
	"strings"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

func (c *Controller) RefreshToken(ctx *gin.Context) {
	var body RefreshTokenRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	body.RefreshToken = strings.TrimSpace(body.RefreshToken)
	if body.RefreshToken == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	tokens, err := c.svc.RefreshToken(ctx.Request.Context(), body.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidToken):
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		default:
			log.Errf("auth.refresh.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, tokens)
}
