package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PointSetting struct {
	bun.BaseModel     `bun:"table:point_settings,alias:ps"`
	ID                uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	EarnRateAmount    *float64  `bun:"earn_rate_amount"`
	EarnPoints        *int      `bun:"earn_points"`
	MinOrderToEarn    float64   `bun:"min_order_to_earn,notnull,default:0"`
	PointExpiryMonths *int      `bun:"point_expiry_months"`
	IsActive          bool      `bun:"is_active,notnull,default:true"`
	UpdatedAt         time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
