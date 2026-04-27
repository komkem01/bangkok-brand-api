package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShippingRateRule struct {
	bun.BaseModel        `bun:"table:shipping_rate_rules,alias:srr"`
	ID                   uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopShippingMethodID *uuid.UUID `bun:"shop_shipping_method_id,type:uuid"`
	MinWeight            *float64   `bun:"min_weight"`
	MaxWeight            *float64   `bun:"max_weight"`
	MinOrderAmount       *float64   `bun:"min_order_amount"`
	MaxOrderAmount       *float64   `bun:"max_order_amount"`
	RateAmount           float64    `bun:"rate_amount,notnull,default:0"`
	Priority             int        `bun:"priority,notnull,default:1"`
	IsActive             bool       `bun:"is_active,notnull,default:true"`
	CreatedAt            time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt            time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
