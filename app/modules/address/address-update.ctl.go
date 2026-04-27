package address

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBodyRequest struct {
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

// Update godoc
// PATCH /addresses/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.AddressInvalidID, nil)
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	id := uuid.MustParse(uri.ID)
	memberID, err := currentMemberID(ctx)
	if err != nil {
		base.Unauthorized(ctx, i18n.Unauthorized, nil)
		return
	}

	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.AddressNotFound, nil)
			return
		}
		log.Errf("address.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressUpdateFailed, nil)
		return
	}

	if current.MemberID == nil || *current.MemberID != memberID {
		base.BadRequest(ctx, i18n.AddressNotFound, nil)
		return
	}

	input := *current
	if body.AddressTypeID != nil {
		v, err := uuid.Parse(*body.AddressTypeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.AddressTypeID = &v
	}
	if body.ProvinceID != nil {
		v, err := uuid.Parse(*body.ProvinceID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ProvinceID = &v
	}
	if body.DistrictID != nil {
		v, err := uuid.Parse(*body.DistrictID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.DistrictID = &v
	}
	if body.SubDistrictID != nil {
		v, err := uuid.Parse(*body.SubDistrictID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.SubDistrictID = &v
	}
	if body.ZipcodeID != nil {
		v, err := uuid.Parse(*body.ZipcodeID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.ZipcodeID = &v
	}
	if body.AddressName != nil {
		input.AddressName = body.AddressName
	}
	if body.RecipientName != nil {
		input.RecipientName = body.RecipientName
	}
	if body.RecipientPhone != nil {
		input.RecipientPhone = body.RecipientPhone
	}
	if body.AddressDetail != nil {
		input.AddressDetail = body.AddressDetail
	}
	if body.IsDefault != nil {
		input.IsDefault = *body.IsDefault
	}
	if body.Latitude != nil {
		input.Latitude = body.Latitude
	}
	if body.Longitude != nil {
		input.Longitude = body.Longitude
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("address.update.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
