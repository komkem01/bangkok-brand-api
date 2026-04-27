package category

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	ParentID    *string `json:"parent_id"`
	NameTh      *string `json:"name_th"`
	NameEn      *string `json:"name_en"`
	Description *string `json:"description"`
	ImageID     *string `json:"image_id"`
	Slug        *string `json:"slug"`
	IsActive    *bool   `json:"is_active"`
	SortOrder   *int    `json:"sort_order"`
}

// Create godoc
// POST /categories
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.Category{
		NameTh:      body.NameTh,
		NameEn:      body.NameEn,
		Description: body.Description,
		Slug:        body.Slug,
		IsActive:    true,
		SortOrder:   0,
	}
	if body.ParentID != nil {
		v, err := uuid.Parse(*body.ParentID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ParentID = &v
	}
	if body.ImageID != nil {
		v, err := uuid.Parse(*body.ImageID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ImageID = &v
	}
	if body.IsActive != nil {
		item.IsActive = *body.IsActive
	}
	if body.SortOrder != nil {
		item.SortOrder = *body.SortOrder
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("category.create.error: %v", err)
		base.InternalServerError(ctx, i18n.CategoryCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
