package address

import (
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
// DELETE /addresses/:id
func (c *Controller) Delete(ctx *gin.Context) {
	var req DeleteRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.AddressInvalidID, nil)
		return
	}

	id := uuid.MustParse(req.ID)
	if err := c.svc.Delete(ctx.Request.Context(), id); err != nil {
		log.Errf("address.delete.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressDeleteFailed, nil)
		return
	}

	base.Success(ctx, nil)
}
