package category

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InfoRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (c *Controller) Info(ctx *gin.Context) {
	var req InfoRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.CategoryInvalidID, nil)
		return
	}

	id := uuid.MustParse(req.ID)
	item, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.CategoryNotFound, nil)
			return
		}
		log.Errf("category.info.error: %v", err)
		base.InternalServerError(ctx, i18n.CategoryInfoFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(item))
}
