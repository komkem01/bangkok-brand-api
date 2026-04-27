package subdistrict

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID         string  `json:"id"`
	DistrictID *string `json:"district_id"`
	Name       string  `json:"name"`
	IsActive   bool    `json:"is_active"`
}

func toListResponse(sd *ent.Subdistrict) ListResponse {
	var districtID *string
	if sd.DistrictID != nil {
		s := sd.DistrictID.String()
		districtID = &s
	}
	return ListResponse{
		ID:         sd.ID.String(),
		DistrictID: districtID,
		Name:       sd.Name,
		IsActive:   sd.IsActive,
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	subdistricts, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("subdistrict.list.error: %v", err)
		base.InternalServerError(ctx, i18n.SubdistrictListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(subdistricts))
	for _, sd := range subdistricts {
		res = append(res, toListResponse(sd))
	}

	base.Success(ctx, res)
}
