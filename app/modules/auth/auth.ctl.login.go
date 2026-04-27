package auth

import (
	"errors"
	"strings"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	var body LoginRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if !isValidEmail(strings.TrimSpace(body.Email)) {
		base.BadRequest(ctx, i18n.MemberInvalidEmail, nil)
		return
	}

	result, err := c.svc.Login(ctx.Request.Context(), &LoginInput{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidCredentials):
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		default:
			log.Errf("auth.login.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, result)
}
