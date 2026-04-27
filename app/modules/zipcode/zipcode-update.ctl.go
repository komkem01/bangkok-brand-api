package zipcode

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
	SubDistrictID *string `json:"sub_district_id"`
	Name          string  `json:"name"`
	IsActive      *bool   `json:"is_active"`
}

func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ZipcodeInvalidID, nil)
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
			base.BadRequest(ctx, i18n.ZipcodeNotFound, nil)
			return
		}
		log.Errf("zipcode.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ZipcodeUpdateFailed, nil)
		return
	}

	input := UpdateInput{
		SubDistrictID: current.SubDistrictID,
		Name:          current.Name,
		IsActive:      current.IsActive,
	}
	if body.SubDistrictID != nil {
		sid := uuid.MustParse(*body.SubDistrictID)
		input.SubDistrictID = &sid
	}
	if body.Name != "" {
		input.Name = body.Name
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}

	zipcode, err := c.svc.Update(ctx.Request.Context(), id, input)
	if err != nil {
		log.Errf("zipcode.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ZipcodeUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(zipcode))
}
