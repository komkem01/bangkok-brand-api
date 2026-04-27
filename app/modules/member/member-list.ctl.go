package member

import (
	"time"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID             string  `json:"id"`
	GenderID       *string `json:"gender_id,omitempty"`
	PrefixID       *string `json:"prefix_id,omitempty"`
	Email          *string `json:"email,omitempty"`
	MemberNo       *string `json:"member_no,omitempty"`
	ProfileImageID *string `json:"profile_image_id,omitempty"`
	Displayname    *string `json:"displayname,omitempty"`
	FirstnameTh    *string `json:"firstname_th,omitempty"`
	LastnameTh     *string `json:"lastname_th,omitempty"`
	CitizenID      *string `json:"citizen_id,omitempty"`
	Birthdate      *string `json:"birthdate,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	Role           string  `json:"role"`
	Status         string  `json:"status"`
	ProvinceID     *string `json:"province_id,omitempty"`
	DistrictID     *string `json:"district_id,omitempty"`
	SubDistrictID  *string `json:"sub_district_id,omitempty"`
	ZipcodeID      *string `json:"zipcode_id,omitempty"`
	RegisterdAt    *string `json:"registerd_at,omitempty"`
	LastedLogin    *string `json:"lasted_login,omitempty"`
	IsVerified     bool    `json:"is_verified"`
	TotalPoints    int     `json:"total_points"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func toListResponse(m *ent.Member) ListResponse {
	res := ListResponse{
		ID:          m.ID.String(),
		Email:       m.Email,
		MemberNo:    m.MemberNo,
		Displayname: m.Displayname,
		FirstnameTh: m.FirstnameTh,
		LastnameTh:  m.LastnameTh,
		CitizenID:   m.CitizenID,
		Phone:       m.Phone,
		Role:        string(m.Role),
		Status:      string(m.Status),
		IsVerified:  m.IsVerified,
		TotalPoints: m.TotalPoints,
		CreatedAt:   m.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:   m.UpdatedAt.Format(utils.RFC3339Milli),
	}
	res.GenderID = uuidToStringPtr(m.GenderID)
	res.PrefixID = uuidToStringPtr(m.PrefixID)
	res.ProfileImageID = uuidToStringPtr(m.ProfileImageID)
	res.ProvinceID = uuidToStringPtr(m.ProvinceID)
	res.DistrictID = uuidToStringPtr(m.DistrictID)
	res.SubDistrictID = uuidToStringPtr(m.SubDistrictID)
	res.ZipcodeID = uuidToStringPtr(m.ZipcodeID)
	res.Birthdate = dateToStringPtr(m.Birthdate)
	res.RegisterdAt = timeToStringPtr(m.RegisterdAt)
	res.LastedLogin = timeToStringPtr(m.LastedLogin)
	return res
}

func uuidToStringPtr(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	s := id.String()
	return &s
}

func timeToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format(utils.RFC3339Milli)
	return &s
}

func dateToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02")
	return &s
}

// List godoc
// GET /members
func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	members, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("member.list.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(members))
	for _, m := range members {
		res = append(res, toListResponse(m))
	}

	base.Success(ctx, res)
}
