package bank

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type CreateBodyRequest struct {
	NameTh   string `json:"name_th"`
	NameEn   string `json:"name_en"`
	Code     string `json:"code"`
	IsActive *bool  `json:"is_active"`
}

// Create godoc
// POST /systems/banks
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.Bank{
		NameTh:   body.NameTh,
		NameEn:   body.NameEn,
		Code:     body.Code,
		IsActive: true,
	}
	if body.IsActive != nil {
		item.IsActive = *body.IsActive
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("bank.create.error: %v", err)
		base.InternalServerError(ctx, i18n.BankCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
