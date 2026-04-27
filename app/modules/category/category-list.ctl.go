package category

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID          string  `json:"id"`
	ParentID    *string `json:"parent_id,omitempty"`
	NameTh      *string `json:"name_th,omitempty"`
	NameEn      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
	ImageID     *string `json:"image_id,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	IsActive    bool    `json:"is_active"`
	SortOrder   int     `json:"sort_order"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func toListResponse(c *ent.Category) ListResponse {
	return ListResponse{
		ID:          c.ID.String(),
		ParentID:    uuidToStringPtr(c.ParentID),
		NameTh:      c.NameTh,
		NameEn:      c.NameEn,
		Description: c.Description,
		ImageID:     uuidToStringPtr(c.ImageID),
		Slug:        c.Slug,
		IsActive:    c.IsActive,
		SortOrder:   c.SortOrder,
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

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("category.list.error: %v", err)
		base.InternalServerError(ctx, i18n.CategoryListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
