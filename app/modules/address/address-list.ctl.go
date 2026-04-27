package address

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID             string   `json:"id"`
	MemberID       *string  `json:"member_id,omitempty"`
	AddressTypeID  *string  `json:"address_type_id,omitempty"`
	AddressName    *string  `json:"address_name,omitempty"`
	RecipientName  *string  `json:"recipient_name,omitempty"`
	RecipientPhone *string  `json:"recipient_phone,omitempty"`
	AddressDetail  *string  `json:"address_detail,omitempty"`
	ProvinceID     *string  `json:"province_id,omitempty"`
	DistrictID     *string  `json:"district_id,omitempty"`
	SubDistrictID  *string  `json:"sub_district_id,omitempty"`
	ZipcodeID      *string  `json:"zipcode_id,omitempty"`
	IsDefault      bool     `json:"is_default"`
	Latitude       *float64 `json:"latitude,omitempty"`
	Longitude      *float64 `json:"longitude,omitempty"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

func toListResponse(a *ent.MemberAddress) ListResponse {
	return ListResponse{
		ID:             a.ID.String(),
		MemberID:       uuidToStringPtr(a.MemberID),
		AddressTypeID:  uuidToStringPtr(a.AddressTypeID),
		AddressName:    a.AddressName,
		RecipientName:  a.RecipientName,
		RecipientPhone: a.RecipientPhone,
		AddressDetail:  a.AddressDetail,
		ProvinceID:     uuidToStringPtr(a.ProvinceID),
		DistrictID:     uuidToStringPtr(a.DistrictID),
		SubDistrictID:  uuidToStringPtr(a.SubDistrictID),
		ZipcodeID:      uuidToStringPtr(a.ZipcodeID),
		IsDefault:      a.IsDefault,
		Latitude:       a.Latitude,
		Longitude:      a.Longitude,
		CreatedAt:      a.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:      a.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func uuidToStringPtr(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	s := id.String()
	return &s
}

// List godoc
// GET /addresses
func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("address.list.error: %v", err)
		base.InternalServerError(ctx, i18n.AddressListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
