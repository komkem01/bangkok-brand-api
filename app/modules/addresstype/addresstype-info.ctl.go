package addresstype

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
	var uri InfoUriRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.AddressTypeInvalidID, nil)
		return
	}

	item, err := c.svc.Info(ctx.Request.Context(), uuid.MustParse(uri.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.AddressTypeNotFound, nil)
			return
		}
		log.Errf("addresstype.info.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressTypeInfoFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(item))
}
