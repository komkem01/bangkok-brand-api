package brand

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
	NameTh      *string `json:"name_th"`
	NameEn      *string `json:"name_en"`
	LogoID      *string `json:"logo_id"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

// Update godoc
// PATCH /brands/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.BrandInvalidID, nil)
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
			base.BadRequest(ctx, i18n.BrandNotFound, nil)
			return
		}
		log.Errf("brand.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.BrandUpdateFailed, nil)
		return
	}

	input := *current
	if body.NameTh != nil {
		input.NameTh = body.NameTh
	}
	if body.NameEn != nil {
		input.NameEn = body.NameEn
	}
	if body.LogoID != nil {
		v, err := uuid.Parse(*body.LogoID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.LogoID = &v
	}
	if body.Description != nil {
		input.Description = body.Description
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("brand.update.error: %v", err)
		base.InternalServerError(ctx, i18n.BrandUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
