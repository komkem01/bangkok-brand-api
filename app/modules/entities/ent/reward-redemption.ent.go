package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RewardRedemption struct {
	bun.BaseModel   `bun:"table:reward_redemptions,alias:rrd"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID        *uuid.UUID `bun:"member_id,type:uuid"`
	RewardID        *uuid.UUID `bun:"reward_id,type:uuid"`
	PointsUsed      *int       `bun:"points_used"`
	Status          string     `bun:"status,notnull,default:pending"`
	RecipientName   *string    `bun:"recipient_name"`
	RecipientPhone  *string    `bun:"recipient_phone"`
	ShippingAddress *string    `bun:"shipping_address"`
	RedeemedAt      *time.Time `bun:"redeemed_at"`
	CreatedAt       time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
