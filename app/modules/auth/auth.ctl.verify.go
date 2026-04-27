package auth

import (
	"errors"
	"strings"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Verify(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)
	defer span.End()

	token, err := extractBearerToken(ctx.GetHeader("Authorization"))
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	member, err := c.svc.VerifyAccessToken(ctx.Request.Context(), token)
	if err != nil {
		switch {
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		default:
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		}
		return
	}

	base.Success(ctx, toAuthMember(member))
}

func extractBearerToken(authorization string) (string, error) {
	parts := strings.Fields(strings.TrimSpace(authorization))
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
		return "", ErrInvalidToken
	}

	return parts[1], nil
}
