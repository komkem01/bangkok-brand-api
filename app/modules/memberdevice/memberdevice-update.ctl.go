package memberdevice

import (
	"database/sql"
	"errors"
	"strings"

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
	Platform    *string `json:"platform"`
	DeviceToken *string `json:"device_token"`
	DeviceName  *string `json:"device_name"`
	AppVersion  *string `json:"app_version"`
	IsActive    *bool   `json:"is_active"`
}

// Update godoc
// PATCH /memberdevices/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.MemberDeviceInvalidID, nil)
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	id := uuid.MustParse(uri.ID)
	memberID, err := currentMemberID(ctx)
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.MemberDeviceNotFound, nil)
			return
		}
		log.Errf("memberdevice.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberDeviceUpdateFailed, nil)
		return
	}

	if current.MemberID == nil || *current.MemberID != memberID {
		base.BadRequest(ctx, i18n.MemberDeviceNotFound, nil)
		return
	}

	input := *current
	if body.Platform != nil {
		v := strings.TrimSpace(*body.Platform)
		if v == "" {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.Platform = v
	}
	if body.DeviceToken != nil {
		v := strings.TrimSpace(*body.DeviceToken)
		if v == "" {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.DeviceToken = v
	}
	if body.DeviceName != nil {
		input.DeviceName = body.DeviceName
	}
	if body.AppVersion != nil {
		input.AppVersion = body.AppVersion
	}
	if body.IsActive != nil {
		input.IsActive = *body.IsActive
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("memberdevice.update.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberDeviceUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
