package brand

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	NameTh      *string `json:"name_th"`
	NameEn      *string `json:"name_en"`
	LogoID      *string `json:"logo_id"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

// Create godoc
// POST /brands
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.Brand{
		NameTh:      body.NameTh,
		NameEn:      body.NameEn,
		Description: body.Description,
		IsActive:    true,
	}
	if body.LogoID != nil {
		v, err := uuid.Parse(*body.LogoID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.LogoID = &v
	}
	if body.IsActive != nil {
		item.IsActive = *body.IsActive
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("brand.create.error: %v", err)
		base.InternalServerError(ctx, i18n.BrandCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
