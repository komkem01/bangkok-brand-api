package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CouponUsage struct {
	bun.BaseModel   `bun:"table:coupon_usages,alias:cu"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CouponID        *uuid.UUID `bun:"coupon_id,type:uuid"`
	MemberID        *uuid.UUID `bun:"member_id,type:uuid"`
	OrderID         *uuid.UUID `bun:"order_id,type:uuid"`
	DiscountApplied *float64   `bun:"discount_applied"`
	CreatedAt       time.Time  `bun:"created_at"`
}
