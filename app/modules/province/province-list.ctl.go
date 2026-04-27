package province

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func toListResponse(p *ent.Province) ListResponse {
	return ListResponse{
		ID:       p.ID.String(),
		Name:     p.Name,
		IsActive: p.IsActive,
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	provinces, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("province.list.error: %v", err)
		base.InternalServerError(ctx, "province-list-failed", nil)
		return
	}

	res := make([]ListResponse, 0, len(provinces))
	for _, p := range provinces {
		res = append(res, toListResponse(p))
	}

	base.Success(ctx, res)
}
