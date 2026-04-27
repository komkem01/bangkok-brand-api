package prefix

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBodyRequest struct {
	GenderID *string `json:"gender_id"`
	NameTh   string  `json:"name_th"`
	NameEn   string  `json:"name_en"`
	IsActive *bool   `json:"is_active"`
}

// Update godoc
// PATCH /prefixes/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	id := uuid.MustParse(uri.ID)

	// Fetch current values for PATCH semantics
	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, "prefix-not-found", nil)
			return
		}
		log.Errf("prefix.update.fetch.error: %v", err)
		base.InternalServerError(ctx, "prefix-update-failed", nil)
		return
	}

	input := UpdateInput{
		GenderID: current.GenderID,
		NameTh:   current.NameTh,
		NameEn:   current.NameEn,
		IsActive: current.IsActive,
	}
	if body.GenderID != nil {
		gid, err := uuid.Parse(*body.GenderID)
		if err != nil {
			base.BadRequest(ctx, "invalid-gender-id", nil)
			return
		}
		input.GenderID = &gid
	}
	if body.NameTh != "" {
		input.NameTh = body.NameTh
	}
	if body.NameEn != "" {
		input.NameEn = body.NameEn
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}

	p, err := c.svc.Update(ctx.Request.Context(), id, input)
	if err != nil {
		log.Errf("prefix.update.error: %v", err)
		base.InternalServerError(ctx, "prefix-update-failed", nil)
		return
	}

	base.Success(ctx, toListResponse(p))
}
