package addresstype

import (
	"strings"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type CreateBodyRequest struct {
	NameTh   string `json:"name_th"`
	NameEn   string `json:"name_en"`
	IsActive *bool  `json:"is_active"`
}

func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if strings.TrimSpace(body.NameTh) == "" || strings.TrimSpace(body.NameEn) == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.AddressType{
		NameTh:   strings.TrimSpace(body.NameTh),
		NameEn:   strings.TrimSpace(body.NameEn),
		IsActive: true,
	}
	if body.IsActive != nil {
		item.IsActive = *body.IsActive
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("addresstype.create.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressTypeCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
