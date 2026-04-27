package addresstype

import (
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (c *Controller) Delete(ctx *gin.Context) {
	var uri DeleteUriRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.AddressTypeInvalidID, nil)
		return
	}

	if err := c.svc.Delete(ctx.Request.Context(), uuid.MustParse(uri.ID)); err != nil {
		log.Errf("addresstype.delete.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressTypeDeleteFailed, nil)
		return
	}

	base.Success(ctx, nil)
}
