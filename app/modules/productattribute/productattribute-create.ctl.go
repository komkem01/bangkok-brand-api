package productattribute

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type CreateBodyRequest struct {
	NameTh *string `json:"name_th"`
	NameEn *string `json:"name_en"`
}

// Create godoc
// POST /product-attributes
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.ProductAttribute{NameTh: body.NameTh, NameEn: body.NameEn}
	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("product-attribute.create.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
