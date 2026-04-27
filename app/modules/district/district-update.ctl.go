package district

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
	ProvinceID *string `json:"province_id"`
	Name       string  `json:"name"`
	IsActive   *bool   `json:"is_active"`
}

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
	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, "district-not-found", nil)
			return
		}
		log.Errf("district.update.fetch.error: %v", err)
		base.InternalServerError(ctx, "district-update-failed", nil)
		return
	}

	input := UpdateInput{
		ProvinceID: current.ProvinceID,
		Name:       current.Name,
		IsActive:   current.IsActive,
	}
	if body.ProvinceID != nil {
		pid := uuid.MustParse(*body.ProvinceID)
		input.ProvinceID = &pid
	}
	if body.Name != "" {
		input.Name = body.Name
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}

	district, err := c.svc.Update(ctx.Request.Context(), id, input)
	if err != nil {
		log.Errf("district.update.error: %v", err)
		base.InternalServerError(ctx, "district-update-failed", nil)
		return
	}

	base.Success(ctx, toListResponse(district))
}
