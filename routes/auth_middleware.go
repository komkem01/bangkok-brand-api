package routes

import (
	"errors"
	"strings"

	"bangkok-brand/app/modules"
	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

func authRequired(mod *modules.Modules) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractBearerToken(ctx.GetHeader("Authorization"))
		if err != nil {
			_ = base.Unauthorized(ctx, i18n.Unauthorized, nil)
			ctx.Abort()
			return
		}

		member, err := mod.Auth.Svc.VerifyAccessToken(ctx.Request.Context(), token)
		if err != nil {
			switch {
			case errors.Is(err, auth.ErrMemberNotActive):
				_ = base.Forbidden(ctx, i18n.Forbidden, nil)
			default:
				_ = base.Unauthorized(ctx, i18n.Unauthorized, nil)
			}
			ctx.Abort()
			return
		}

		ctx.Set(auth.ContextKeyMemberID, member.ID)
		ctx.Set(auth.ContextKeyMember, auth.AuthMember{
			ID:          member.ID,
			Email:       member.Email,
			Phone:       member.Phone,
			Firstname:   member.FirstnameTh,
			Lastname:    member.LastnameTh,
			Role:        member.Role,
			Status:      member.Status,
			IsVerified:  member.IsVerified,
			LastedLogin: member.LastedLogin,
		})

		ctx.Next()
	}
}

func extractBearerToken(authorization string) (string, error) {
	parts := strings.Fields(strings.TrimSpace(authorization))
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
		return "", auth.ErrInvalidToken
	}

	return parts[1], nil
}
