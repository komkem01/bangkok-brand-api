package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Order represents a row in orders.
type Order struct {
	bun.BaseModel `bun:"table:orders,alias:o"`

	ID                    uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrderNo               *string    `bun:"order_no"`
	MemberID              *uuid.UUID `bun:"member_id,type:uuid"`
	TotalProductPrice     *float64   `bun:"total_product_price"`
	ShippingFee           float64    `bun:"shipping_fee,notnull,default:0"`
	DiscountAmount        float64    `bun:"discount_amount,notnull,default:0"`
	NetAmount             *float64   `bun:"net_amount"`
	RecipientName         *string    `bun:"recipient_name"`
	RecipientPhone        *string    `bun:"recipient_phone"`
	ShippingAddressDetail *string    `bun:"shipping_address_detail"`
	ProvinceID            *uuid.UUID `bun:"province_id,type:uuid"`
	DistrictID            *uuid.UUID `bun:"district_id,type:uuid"`
	SubDistrictID         *uuid.UUID `bun:"sub_district_id,type:uuid"`
	ZipcodeID             *uuid.UUID `bun:"zipcode_id,type:uuid"`
	Status                string     `bun:"status,notnull,default:pending_payment"`
	TrackingNumber        *string    `bun:"tracking_number"`
	CourierName           *string    `bun:"courier_name"`
	Remark                *string    `bun:"remark"`
	OrderedAt             *time.Time `bun:"ordered_at"`
	CreatedAt             time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt             time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
