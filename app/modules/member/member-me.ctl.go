package member

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// InfoMe godoc
// GET /members/me
func (c *Controller) InfoMe(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	memberIDRaw, ok := ctx.Get(auth.ContextKeyMemberID)
	if !ok {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	var id uuid.UUID
	switch v := memberIDRaw.(type) {
	case uuid.UUID:
		id = v
	case string:
		parsed, err := uuid.Parse(v)
		if err != nil {
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
			return
		}
		id = parsed
	default:
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	m, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.MemberNotFound, nil)
			return
		}
		log.Errf("member.me.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberInfoFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(m))
}
