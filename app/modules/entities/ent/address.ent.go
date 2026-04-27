package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// MemberAddress represents a row in member_addresses.
type MemberAddress struct {
	bun.BaseModel `bun:"table:member_addresses,alias:ma"`

	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID       *uuid.UUID `bun:"member_id,type:uuid"`
	AddressTypeID  *uuid.UUID `bun:"address_type_id,type:uuid"`
	AddressName    *string    `bun:"address_name"`
	RecipientName  *string    `bun:"recipient_name"`
	RecipientPhone *string    `bun:"recipient_phone"`
	AddressDetail  *string    `bun:"address_detail"`
	ProvinceID     *uuid.UUID `bun:"province_id,type:uuid"`
	DistrictID     *uuid.UUID `bun:"district_id,type:uuid"`
	SubDistrictID  *uuid.UUID `bun:"sub_district_id,type:uuid"`
	ZipcodeID      *uuid.UUID `bun:"zipcode_id,type:uuid"`
	IsDefault      bool       `bun:"is_default,notnull,default:false"`
	Latitude       *float64   `bun:"latitude"`
	Longitude      *float64   `bun:"longitude"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
