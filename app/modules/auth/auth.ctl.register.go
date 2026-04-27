package auth

import (
	"errors"
	"strings"
	"time"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) Register(ctx *gin.Context) {
	var body RegisterRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if !isValidEmail(strings.TrimSpace(body.Email)) {
		base.BadRequest(ctx, i18n.MemberInvalidEmail, nil)
		return
	}

	genderID, err := uuid.Parse(strings.TrimSpace(body.GenderID))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	prefixID, err := uuid.Parse(strings.TrimSpace(body.PrefixID))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	provinceID, err := uuid.Parse(strings.TrimSpace(body.ProvinceID))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	districtID, err := uuid.Parse(strings.TrimSpace(body.DistrictID))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	subDistrictID, err := uuid.Parse(strings.TrimSpace(body.SubDistrictID))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	zipcodeID, err := uuid.Parse(strings.TrimSpace(body.ZipcodeID))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	citizenID := strings.TrimSpace(body.CitizenID)
	if !registerCitizenIDPattern.MatchString(citizenID) {
		base.BadRequest(ctx, i18n.MemberInvalidCitizenID, nil)
		return
	}

	phone := strings.TrimSpace(body.Phone)
	if !registerPhonePattern.MatchString(phone) {
		base.BadRequest(ctx, i18n.MemberInvalidPhone, nil)
		return
	}

	birthdate, err := time.Parse("2006-01-02", strings.TrimSpace(body.Birthdate))
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	displayname := strings.TrimSpace(body.Displayname)
	firstname := strings.TrimSpace(body.FirstnameTh)
	lastname := strings.TrimSpace(body.LastnameTh)
	if displayname == "" || firstname == "" || lastname == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	result, err := c.svc.Register(ctx.Request.Context(), &RegisterInput{
		GenderID:      &genderID,
		PrefixID:      &prefixID,
		Email:         body.Email,
		Password:      body.Password,
		Displayname:   &displayname,
		Firstname:     &firstname,
		Lastname:      &lastname,
		CitizenID:     &citizenID,
		Birthdate:     &birthdate,
		Phone:         &phone,
		ProvinceID:    &provinceID,
		DistrictID:    &districtID,
		SubDistrictID: &subDistrictID,
		ZipcodeID:     &zipcodeID,
	})
	if err != nil {
		switch {
		case errors.Is(err, ErrEmailExists):
			base.BadRequest(ctx, i18n.BadRequest, nil)
		case errors.Is(err, ErrInvalidCredentials):
			base.BadRequest(ctx, i18n.BadRequest, nil)
		default:
			log.Errf("auth.register.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, result)
}
