package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Cart represents a row in carts.
type Cart struct {
	bun.BaseModel `bun:"table:carts,alias:c"`

	ID         uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID   *uuid.UUID `bun:"member_id,type:uuid"`
	TotalItems int        `bun:"total_items,notnull,default:0"`
	TotalPrice float64    `bun:"total_price,notnull,default:0"`
	CreatedAt  time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt  time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
