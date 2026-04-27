package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Coupon struct {
	bun.BaseModel     `bun:"table:coupons,alias:co"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Code              *string    `bun:"code"`
	NameTh            *string    `bun:"name_th"`
	Description       *string    `bun:"description"`
	Type              *string    `bun:"type"`
	Value             *float64   `bun:"value"`
	MaxDiscountAmount *float64   `bun:"max_discount_amount"`
	MinOrderAmount    float64    `bun:"min_order_amount"`
	LimitPerCoupon    *int       `bun:"limit_per_coupon"`
	LimitPerMember    int        `bun:"limit_per_member"`
	UsedCount         int        `bun:"used_count"`
	StartDate         *time.Time `bun:"start_date"`
	EndDate           *time.Time `bun:"end_date"`
	IsActive          bool       `bun:"is_active"`
	CreatedAt         time.Time  `bun:"created_at"`
	UpdatedAt         time.Time  `bun:"updated_at"`
}
