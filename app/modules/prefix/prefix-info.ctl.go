package prefix

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InfoRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// Info godoc
// GET /prefixes/:id
func (c *Controller) Info(ctx *gin.Context) {
	var req InfoRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	id := uuid.MustParse(req.ID)
	p, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, "prefix-not-found", nil)
			return
		}
		log.Errf("prefix.info.error: %v", err)
		base.InternalServerError(ctx, "prefix-info-failed", nil)
		return
	}

	base.Success(ctx, toListResponse(p))
}
