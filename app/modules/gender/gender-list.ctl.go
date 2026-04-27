package gender

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID       string `json:"id"`
	NameTh   string `json:"name_th"`
	NameEn   string `json:"name_en"`
	IsActive bool   `json:"is_active"`
}

func toListResponse(g *ent.Gender) ListResponse {
	return ListResponse{
		ID:       g.ID.String(),
		NameTh:   g.NameTh,
		NameEn:   g.NameEn,
		IsActive: g.IsActive,
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	genders, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("gender.list.error: %v", err)
		base.InternalServerError(ctx, "gender-list-failed", nil)
		return
	}

	res := make([]ListResponse, 0, len(genders))
	for _, g := range genders {
		res = append(res, toListResponse(g))
	}

	base.Success(ctx, res)
}
