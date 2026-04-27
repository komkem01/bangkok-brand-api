package province

import (
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"

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
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	id := uuid.MustParse(req.ID)
	if err := c.svc.Delete(ctx.Request.Context(), id); err != nil {
		log.Errf("province.delete.error: %v", err)
		base.InternalServerError(ctx, "province-delete-failed", nil)
		return
	}

	base.Success(ctx, nil)
}
