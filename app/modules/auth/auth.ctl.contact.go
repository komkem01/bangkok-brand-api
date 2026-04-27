package auth

import (
	"errors"
	"strings"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type UpdateContactInfoRequest struct {
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

type ContactInfoResponse struct {
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

func (c *Controller) ContactInfo(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	token, err := extractBearerToken(ctx.GetHeader("Authorization"))
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	member, err := c.svc.ContactInfo(ctx.Request.Context(), token)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidToken):
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		default:
			log.Errf("auth.contact.info.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, ContactInfoResponse{
		Email: member.Email,
		Phone: member.Phone,
	})
}

func (c *Controller) UpdateContactInfo(ctx *gin.Context) {
	var body UpdateContactInfoRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if body.Email == nil && body.Phone == nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	var email *string
	if body.Email != nil {
		trimmed := strings.TrimSpace(*body.Email)
		if !isValidEmail(trimmed) {
			base.BadRequest(ctx, i18n.MemberInvalidEmail, nil)
			return
		}
		email = &trimmed
	}

	var phone *string
	if body.Phone != nil {
		trimmed := strings.TrimSpace(*body.Phone)
		if !registerPhonePattern.MatchString(trimmed) {
			base.BadRequest(ctx, i18n.MemberInvalidPhone, nil)
			return
		}
		phone = &trimmed
	}

	token, err := extractBearerToken(ctx.GetHeader("Authorization"))
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	member, err := c.svc.UpdateContactInfo(ctx.Request.Context(), token, email, phone)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidToken):
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		case errors.Is(err, ErrEmailExists):
			base.BadRequest(ctx, i18n.BadRequest, nil)
		default:
			log.Errf("auth.contact.update.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, ContactInfoResponse{
		Email: member.Email,
		Phone: member.Phone,
	})
}
