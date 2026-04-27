package member

import (
	"database/sql"
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	citizenIDPattern = regexp.MustCompile(`^\d{13}$`)
	phonePattern     = regexp.MustCompile(`^\+?[0-9]{9,15}$`)
)

type UpdateUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBodyRequest struct {
	GenderID       *string `json:"gender_id"`
	PrefixID       *string `json:"prefix_id"`
	Email          *string `json:"email"`
	Password       *string `json:"password"`
	ProfileImageID *string `json:"profile_image_id"`
	Displayname    *string `json:"displayname"`
	FirstnameTh    *string `json:"firstname_th"`
	LastnameTh     *string `json:"lastname_th"`
	CitizenID      *string `json:"citizen_id"`
	Birthdate      *string `json:"birthdate"`
	Phone          *string `json:"phone"`
	Role           *string `json:"role"`
	Status         *string `json:"status"`
	ProvinceID     *string `json:"province_id"`
	DistrictID     *string `json:"district_id"`
	SubDistrictID  *string `json:"sub_district_id"`
	ZipcodeID      *string `json:"zipcode_id"`
	RegisterdAt    *string `json:"registerd_at"`
	LastedLogin    *string `json:"lasted_login"`
	IsVerified     *bool   `json:"is_verified"`
	TotalPoints    *int    `json:"total_points"`
}

// Update godoc
// PATCH /members/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.MemberInvalidID, nil)
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
			base.BadRequest(ctx, i18n.MemberNotFound, nil)
			return
		}
		log.Errf("member.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberUpdateFailed, nil)
		return
	}

	input := *current
	if body.GenderID != nil {
		val, err := uuid.Parse(*body.GenderID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.GenderID = &val
	}
	if body.PrefixID != nil {
		val, err := uuid.Parse(*body.PrefixID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.PrefixID = &val
	}
	if body.ProfileImageID != nil {
		val, err := uuid.Parse(*body.ProfileImageID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ProfileImageID = &val
	}
	if body.ProvinceID != nil {
		val, err := uuid.Parse(*body.ProvinceID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ProvinceID = &val
	}
	if body.DistrictID != nil {
		val, err := uuid.Parse(*body.DistrictID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.DistrictID = &val
	}
	if body.SubDistrictID != nil {
		val, err := uuid.Parse(*body.SubDistrictID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.SubDistrictID = &val
	}
	if body.ZipcodeID != nil {
		val, err := uuid.Parse(*body.ZipcodeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ZipcodeID = &val
	}

	if body.Email != nil {
		email := strings.TrimSpace(*body.Email)
		if !isValidEmail(email) {
			base.BadRequest(ctx, i18n.MemberInvalidEmail, nil)
			return
		}
		input.Email = &email
	}
	if body.Password != nil {
		input.Password = body.Password
	}
	if body.Displayname != nil {
		input.Displayname = body.Displayname
	}
	if body.FirstnameTh != nil {
		input.FirstnameTh = body.FirstnameTh
	}
	if body.LastnameTh != nil {
		input.LastnameTh = body.LastnameTh
	}
	if body.CitizenID != nil {
		citizenID := strings.TrimSpace(*body.CitizenID)
		if !citizenIDPattern.MatchString(citizenID) {
			base.BadRequest(ctx, i18n.MemberInvalidCitizenID, nil)
			return
		}
		input.CitizenID = &citizenID
	}
	if body.Phone != nil {
		phone := strings.TrimSpace(*body.Phone)
		if !phonePattern.MatchString(phone) {
			base.BadRequest(ctx, i18n.MemberInvalidPhone, nil)
			return
		}
		input.Phone = &phone
	}
	if body.Role != nil {
		role := strings.TrimSpace(*body.Role)
		if !isValidRole(role) {
			base.BadRequest(ctx, i18n.MemberInvalidRole, nil)
			return
		}
		input.Role = ent.MemberRole(role)
	}
	if body.Status != nil {
		status := strings.TrimSpace(*body.Status)
		if !isValidStatus(status) {
			base.BadRequest(ctx, i18n.MemberInvalidStatus, nil)
			return
		}
		input.Status = ent.MemberStatus(status)
	}
	if body.IsVerified != nil {
		input.IsVerified = *body.IsVerified
	}
	if body.TotalPoints != nil {
		input.TotalPoints = *body.TotalPoints
	}
	if body.Birthdate != nil {
		v, err := time.Parse("2006-01-02", *body.Birthdate)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.Birthdate = &v
	}
	if body.RegisterdAt != nil {
		v, err := time.Parse(time.RFC3339, *body.RegisterdAt)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.RegisterdAt = &v
	}
	if body.LastedLogin != nil {
		v, err := time.Parse(time.RFC3339, *body.LastedLogin)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.LastedLogin = &v
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("member.update.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}

func isValidEmail(v string) bool {
	if v == "" {
		return false
	}
	addr, err := mail.ParseAddress(v)
	if err != nil {
		return false
	}
	return addr.Address == v
}

func isValidRole(v string) bool {
	switch ent.MemberRole(v) {
	case ent.MemberRoleCustomer, ent.MemberRoleAdmin, ent.MemberRoleMerchant:
		return true
	default:
		return false
	}
}

func isValidStatus(v string) bool {
	switch ent.MemberStatus(v) {
	case ent.MemberStatusActive, ent.MemberStatusInactive, ent.MemberStatusSuspended:
		return true
	default:
		return false
	}
}
