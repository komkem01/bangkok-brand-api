package memberdevice

import (
	"fmt"
	"time"

	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID          string  `json:"id"`
	MemberID    *string `json:"member_id,omitempty"`
	Platform    string  `json:"platform"`
	DeviceToken string  `json:"device_token"`
	DeviceName  *string `json:"device_name,omitempty"`
	AppVersion  *string `json:"app_version,omitempty"`
	IsActive    bool    `json:"is_active"`
	LastSeenAt  *string `json:"last_seen_at,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func toListResponse(c *ent.MemberDevice) ListResponse {
	var lastSeenAt *string
	if c.LastSeenAt != nil {
		v := c.LastSeenAt.Format(time.RFC3339)
		lastSeenAt = &v
	}

	return ListResponse{
		ID:          c.ID.String(),
		MemberID:    uuidToStringPtr(c.MemberID),
		Platform:    c.Platform,
		DeviceToken: c.DeviceToken,
		DeviceName:  c.DeviceName,
		AppVersion:  c.AppVersion,
		IsActive:    c.IsActive,
		LastSeenAt:  lastSeenAt,
		CreatedAt:   c.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:   c.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func uuidToStringPtr(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	s := id.String()
	return &s
}

// List godoc
// GET /memberdevices
func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	memberID, err := currentMemberID(ctx)
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	items, err := c.svc.ListByMemberID(ctx.Request.Context(), memberID)
	if err != nil {
		log.Errf("memberdevice.list.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberDeviceListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}

func currentMemberID(ctx *gin.Context) (uuid.UUID, error) {
	v, ok := ctx.Get(auth.ContextKeyMemberID)
	if !ok {
		return uuid.Nil, fmt.Errorf("missing member id")
	}

	switch id := v.(type) {
	case uuid.UUID:
		return id, nil
	case string:
		parsed, err := uuid.Parse(id)
		if err != nil {
			return uuid.Nil, err
		}
		return parsed, nil
	default:
		return uuid.Nil, fmt.Errorf("invalid member id type")
	}
}
