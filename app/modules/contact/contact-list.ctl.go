package contact

import (
	"fmt"

	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID            string  `json:"id"`
	MemberID      *string `json:"member_id,omitempty"`
	ContactTypeID *string `json:"contact_type_id,omitempty"`
	Value         *string `json:"value,omitempty"`
	IsPrimary     bool    `json:"is_primary"`
	IsVerified    bool    `json:"is_verified"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func toListResponse(c *ent.MemberContact) ListResponse {
	return ListResponse{
		ID:            c.ID.String(),
		MemberID:      uuidToStringPtr(c.MemberID),
		ContactTypeID: uuidToStringPtr(c.ContactTypeID),
		Value:         c.Value,
		IsPrimary:     c.IsPrimary,
		IsVerified:    c.IsVerified,
		CreatedAt:     c.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:     c.UpdatedAt.Format(utils.RFC3339Milli),
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
// GET /contacts
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
		log.Errf("contact.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ContactListFailed, nil)
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
