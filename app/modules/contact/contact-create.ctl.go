package contact

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	MemberID      *string `json:"member_id"`
	ContactTypeID *string `json:"contact_type_id"`
	Value         *string `json:"value"`
	IsPrimary     *bool   `json:"is_primary"`
	IsVerified    *bool   `json:"is_verified"`
}

// Create godoc
// POST /contacts
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.MemberContact{}
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.MemberID = &v
	}
	if body.ContactTypeID != nil {
		v, err := uuid.Parse(*body.ContactTypeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ContactTypeID = &v
	}
	item.Value = body.Value
	if body.IsPrimary != nil {
		item.IsPrimary = *body.IsPrimary
	}
	if body.IsVerified != nil {
		item.IsVerified = *body.IsVerified
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("contact.create.error: %v", err)
		base.InternalServerError(ctx, i18n.ContactCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
