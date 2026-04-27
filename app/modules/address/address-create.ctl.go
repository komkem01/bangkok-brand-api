package address

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	MemberID       *string  `json:"member_id"`
	AddressTypeID  *string  `json:"address_type_id"`
	AddressName    *string  `json:"address_name"`
	RecipientName  *string  `json:"recipient_name"`
	RecipientPhone *string  `json:"recipient_phone"`
	AddressDetail  *string  `json:"address_detail"`
	ProvinceID     *string  `json:"province_id"`
	DistrictID     *string  `json:"district_id"`
	SubDistrictID  *string  `json:"sub_district_id"`
	ZipcodeID      *string  `json:"zipcode_id"`
	IsDefault      *bool    `json:"is_default"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
}

// Create godoc
// POST /addresses
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.MemberAddress{}
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.MemberID = &v
	}
	if body.AddressTypeID != nil {
		v, err := uuid.Parse(*body.AddressTypeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.AddressTypeID = &v
	}
	if body.ProvinceID != nil {
		v, err := uuid.Parse(*body.ProvinceID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ProvinceID = &v
	}
	if body.DistrictID != nil {
		v, err := uuid.Parse(*body.DistrictID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.DistrictID = &v
	}
	if body.SubDistrictID != nil {
		v, err := uuid.Parse(*body.SubDistrictID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.SubDistrictID = &v
	}
	if body.ZipcodeID != nil {
		v, err := uuid.Parse(*body.ZipcodeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.ZipcodeID = &v
	}
	item.AddressName = body.AddressName
	item.RecipientName = body.RecipientName
	item.RecipientPhone = body.RecipientPhone
	item.AddressDetail = body.AddressDetail
	if body.IsDefault != nil {
		item.IsDefault = *body.IsDefault
	}
	item.Latitude = body.Latitude
	item.Longitude = body.Longitude

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("address.create.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
