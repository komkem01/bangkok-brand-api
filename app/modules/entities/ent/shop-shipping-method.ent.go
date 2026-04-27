package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShopShippingMethod struct {
	bun.BaseModel         `bun:"table:shop_shipping_methods,alias:ssm"`
	ID                    uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID                *uuid.UUID `bun:"shop_id,type:uuid"`
	ShippingMethodID      *uuid.UUID `bun:"shipping_method_id,type:uuid"`
	ShippingZoneID        *uuid.UUID `bun:"shipping_zone_id,type:uuid"`
	MethodName            *string    `bun:"method_name"`
	FeeAdjustment         float64    `bun:"fee_adjustment,notnull,default:0"`
	FreeShippingThreshold *float64   `bun:"free_shipping_threshold"`
	EstimatedDaysMin      *int       `bun:"estimated_days_min"`
	EstimatedDaysMax      *int       `bun:"estimated_days_max"`
	IsCODAvailable        bool       `bun:"is_cod_available,notnull,default:false"`
	IsActive              bool       `bun:"is_active,notnull,default:true"`
	CreatedAt             time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt             time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
