package contact

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
	MemberID      *string `json:"member_id"`
	ContactTypeID *string `json:"contact_type_id"`
	Value         *string `json:"value"`
	IsPrimary     *bool   `json:"is_primary"`
	IsVerified    *bool   `json:"is_verified"`
}

// Update godoc
// PATCH /contacts/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.ContactInvalidID, nil)
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
			base.BadRequest(ctx, i18n.ContactNotFound, nil)
			return
		}
		log.Errf("contact.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.ContactUpdateFailed, nil)
		return
	}

	input := *current
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.MemberID = &v
	}
	if body.ContactTypeID != nil {
		v, err := uuid.Parse(*body.ContactTypeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ContactTypeID = &v
	}
	if body.Value != nil {
		input.Value = body.Value
	}
	if body.IsPrimary != nil {
		input.IsPrimary = *body.IsPrimary
	}
	if body.IsVerified != nil {
		input.IsVerified = *body.IsVerified
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("contact.update.error: %v", err)
		base.InternalServerError(ctx, i18n.ContactUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
