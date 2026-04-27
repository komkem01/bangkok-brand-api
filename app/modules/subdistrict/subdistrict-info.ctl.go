package subdistrict

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InfoUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (c *Controller) Info(ctx *gin.Context) {
	var req InfoUriRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	subdistrict, err := c.svc.Info(ctx.Request.Context(), uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, "subdistrict-not-found", nil)
			return
		}
		log.Errf("subdistrict.info.error: %v", err)
		base.InternalServerError(ctx, "subdistrict-info-failed", nil)
		return
	}

	base.Success(ctx, toListResponse(subdistrict))
}
