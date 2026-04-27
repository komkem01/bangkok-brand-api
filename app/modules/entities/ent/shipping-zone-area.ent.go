package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShippingZoneArea struct {
	bun.BaseModel  `bun:"table:shipping_zone_areas,alias:sza"`
	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShippingZoneID *uuid.UUID `bun:"shipping_zone_id,type:uuid"`
	ProvinceID     *uuid.UUID `bun:"province_id,type:uuid"`
	DistrictID     *uuid.UUID `bun:"district_id,type:uuid"`
	SubDistrictID  *uuid.UUID `bun:"sub_district_id,type:uuid"`
	ZipcodeID      *uuid.UUID `bun:"zipcode_id,type:uuid"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
