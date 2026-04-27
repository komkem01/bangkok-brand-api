package subdistrict

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

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
		base.BadRequest(ctx, i18n.SubdistrictInvalidID, nil)
		return
	}

	subdistrict, err := c.svc.Info(ctx.Request.Context(), uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.SubdistrictNotFound, nil)
			return
		}
		log.Errf("subdistrict.info.error: %v", err)
		base.InternalServerError(ctx, i18n.SubdistrictInfoFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(subdistrict))
}
