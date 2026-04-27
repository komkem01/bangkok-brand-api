package productattribute

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID        string  `json:"id"`
	NameTh    *string `json:"name_th,omitempty"`
	NameEn    *string `json:"name_en,omitempty"`
	CreatedAt string  `json:"created_at"`
}

func toListResponse(p *ent.ProductAttribute) ListResponse {
	return ListResponse{
		ID:        p.ID.String(),
		NameTh:    p.NameTh,
		NameEn:    p.NameEn,
		CreatedAt: p.CreatedAt.Format(utils.RFC3339Milli),
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("product-attribute.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductAttributeListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
