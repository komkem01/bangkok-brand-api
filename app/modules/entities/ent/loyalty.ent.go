package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PointTransaction struct {
	bun.BaseModel   `bun:"table:point_transactions,alias:lo"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID        *uuid.UUID `bun:"member_id,type:uuid"`
	OrderID         *uuid.UUID `bun:"order_id,type:uuid"`
	Type            *string    `bun:"type"`
	Points          *int       `bun:"points"`
	BalanceSnapshot *int       `bun:"balance_snapshot"`
	Description     *string    `bun:"description"`
	ExpiryDate      *time.Time `bun:"expiry_date"`
	CreatedAt       time.Time  `bun:"created_at"`
}
