package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShopSetting struct {
	bun.BaseModel           `bun:"table:shop_settings,alias:ss"`
	ID                      uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID                  *uuid.UUID     `bun:"shop_id,type:uuid"`
	AutoAcceptOrders        bool           `bun:"auto_accept_orders,notnull,default:false"`
	AllowCOD                bool           `bun:"allow_cod,notnull,default:true"`
	MinOrderAmount          float64        `bun:"min_order_amount,notnull,default:0"`
	PreparationTimeMinutes  int            `bun:"preparation_time_minutes,notnull,default:30"`
	DefaultShippingMethodID *uuid.UUID     `bun:"default_shipping_method_id,type:uuid"`
	ReturnPolicy            *string        `bun:"return_policy"`
	RefundPolicy            *string        `bun:"refund_policy"`
	BusinessHours           map[string]any `bun:"business_hours,type:json"`
	CreatedAt               time.Time      `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt               time.Time      `bun:"updated_at,notnull,default:current_timestamp"`
}
