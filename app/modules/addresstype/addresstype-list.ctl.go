package addresstype

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID        string `json:"id"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toListResponse(c *ent.AddressType) ListResponse {
	return ListResponse{
		ID:        c.ID.String(),
		NameTh:    c.NameTh,
		NameEn:    c.NameEn,
		IsActive:  c.IsActive,
		CreatedAt: c.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt: c.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("addresstype.list.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressTypeListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
