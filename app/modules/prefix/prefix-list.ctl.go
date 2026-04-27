package prefix

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID       string  `json:"id"`
	GenderID *string `json:"gender_id,omitempty"`
	NameTh   string  `json:"name_th"`
	NameEn   string  `json:"name_en"`
	IsActive bool    `json:"is_active"`
}

func toListResponse(p *ent.Prefix) ListResponse {
	res := ListResponse{
		ID:       p.ID.String(),
		NameTh:   p.NameTh,
		NameEn:   p.NameEn,
		IsActive: p.IsActive,
	}
	if p.GenderID != nil {
		s := p.GenderID.String()
		res.GenderID = &s
	}
	return res
}

// List godoc
// GET /prefixes
func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	prefixes, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("prefix.list.error: %v", err)
		base.InternalServerError(ctx, i18n.PrefixListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(prefixes))
	for _, p := range prefixes {
		res = append(res, toListResponse(p))
	}

	base.Success(ctx, res)
}
