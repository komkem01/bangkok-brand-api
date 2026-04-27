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

type UpdateUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBodyRequest struct {
	ParentID    *string `json:"parent_id"`
	NameTh      *string `json:"name_th"`
	NameEn      *string `json:"name_en"`
	Description *string `json:"description"`
	ImageID     *string `json:"image_id"`
	Slug        *string `json:"slug"`
	IsActive    *bool   `json:"is_active"`
	SortOrder   *int    `json:"sort_order"`
}

// Update godoc
// PATCH /categories/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.CategoryInvalidID, nil)
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	id := uuid.MustParse(uri.ID)
	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.CategoryNotFound, nil)
			return
		}
		log.Errf("category.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.CategoryUpdateFailed, nil)
		return
	}

	input := *current
	if body.ParentID != nil {
		v, err := uuid.Parse(*body.ParentID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ParentID = &v
	}
	if body.NameTh != nil {
		input.NameTh = body.NameTh
	}
	if body.NameEn != nil {
		input.NameEn = body.NameEn
	}
	if body.Description != nil {
		input.Description = body.Description
	}
	if body.ImageID != nil {
		v, err := uuid.Parse(*body.ImageID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ImageID = &v
	}
	if body.Slug != nil {
		input.Slug = body.Slug
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}
	if body.SortOrder != nil {
		input.SortOrder = *body.SortOrder
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("category.update.error: %v", err)
		base.InternalServerError(ctx, i18n.CategoryUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
