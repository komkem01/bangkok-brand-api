package memberdevice

import (
	"strings"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type CreateBodyRequest struct {
	Platform    string  `json:"platform"`
	DeviceToken string  `json:"device_token"`
	DeviceName  *string `json:"device_name"`
	AppVersion  *string `json:"app_version"`
	IsActive    *bool   `json:"is_active"`
}

// Create godoc
// POST /memberdevices
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if strings.TrimSpace(body.Platform) == "" || strings.TrimSpace(body.DeviceToken) == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	memberID, err := currentMemberID(ctx)
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	item := &ent.MemberDevice{
		MemberID:    &memberID,
		Platform:    strings.TrimSpace(body.Platform),
		DeviceToken: strings.TrimSpace(body.DeviceToken),
		DeviceName:  body.DeviceName,
		AppVersion:  body.AppVersion,
		IsActive:    true,
	}
	if body.IsActive != nil {
		item.IsActive = *body.IsActive
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("memberdevice.create.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberDeviceCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
