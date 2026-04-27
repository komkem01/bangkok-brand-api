package brand

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
	NameTh      *string `json:"name_th,omitempty"`
	NameEn      *string `json:"name_en,omitempty"`
	LogoID      *string `json:"logo_id,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func toListResponse(b *ent.Brand) ListResponse {
	return ListResponse{
		ID:          b.ID.String(),
		NameTh:      b.NameTh,
		NameEn:      b.NameEn,
		LogoID:      uuidToStringPtr(b.LogoID),
		Description: b.Description,
		IsActive:    b.IsActive,
		CreatedAt:   b.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:   b.UpdatedAt.Format(utils.RFC3339Milli),
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
		log.Errf("brand.list.error: %v", err)
		base.InternalServerError(ctx, i18n.BrandListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
