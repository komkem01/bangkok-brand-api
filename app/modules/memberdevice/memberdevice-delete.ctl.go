package memberdevice

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// Delete godoc
// DELETE /memberdevices/:id
func (c *Controller) Delete(ctx *gin.Context) {
	var req DeleteRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.MemberDeviceInvalidID, nil)
		return
	}

	id := uuid.MustParse(req.ID)
	memberID, err := currentMemberID(ctx)
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	item, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.MemberDeviceNotFound, nil)
			return
		}
		log.Errf("memberdevice.delete.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberDeviceDeleteFailed, nil)
		return
	}

	if item.MemberID == nil || *item.MemberID != memberID {
		base.BadRequest(ctx, i18n.MemberDeviceNotFound, nil)
		return
	}

	if err := c.svc.Delete(ctx.Request.Context(), id); err != nil {
		log.Errf("memberdevice.delete.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberDeviceDeleteFailed, nil)
		return
	}

	base.Success(ctx, nil)
}
