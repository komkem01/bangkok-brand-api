package zipcode

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID            string  `json:"id"`
	SubDistrictID *string `json:"sub_district_id"`
	Name          string  `json:"name"`
	IsActive      bool    `json:"is_active"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func toListResponse(z *ent.Zipcode) ListResponse {
	var subDistrictID *string
	if z.SubDistrictID != nil {
		s := z.SubDistrictID.String()
		subDistrictID = &s
	}
	return ListResponse{
		ID:            z.ID.String(),
		SubDistrictID: subDistrictID,
		Name:          z.Name,
		IsActive:      z.IsActive,
		CreatedAt:     z.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:     z.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	zipcodes, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("zipcode.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ZipcodeListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(zipcodes))
	for _, z := range zipcodes {
		res = append(res, toListResponse(z))
	}

	base.Success(ctx, res)
}
