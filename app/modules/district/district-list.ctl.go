package district

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID         string  `json:"id"`
	ProvinceID *string `json:"province_id"`
	Name       string  `json:"name"`
	IsActive   bool    `json:"is_active"`
}

func toListResponse(d *ent.District) ListResponse {
	var provinceID *string
	if d.ProvinceID != nil {
		s := d.ProvinceID.String()
		provinceID = &s
	}
	return ListResponse{
		ID:         d.ID.String(),
		ProvinceID: provinceID,
		Name:       d.Name,
		IsActive:   d.IsActive,
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	districts, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("district.list.error: %v", err)
		base.InternalServerError(ctx, i18n.DistrictListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(districts))
	for _, d := range districts {
		res = append(res, toListResponse(d))
	}

	base.Success(ctx, res)
}
