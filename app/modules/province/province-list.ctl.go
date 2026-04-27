package province

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toListResponse(p *ent.Province) ListResponse {
	return ListResponse{
		ID:        p.ID.String(),
		Name:      p.Name,
		IsActive:  p.IsActive,
		CreatedAt: p.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt: p.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	provinces, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("province.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ProvinceListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(provinces))
	for _, p := range provinces {
		res = append(res, toListResponse(p))
	}

	base.Success(ctx, res)
}
