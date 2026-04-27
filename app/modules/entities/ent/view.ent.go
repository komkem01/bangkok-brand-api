package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductView struct {
	bun.BaseModel    `bun:"table:product_views,alias:vi"`
	ID               uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID         *uuid.UUID `bun:"member_id,type:uuid"`
	SessionID        *string    `bun:"session_id"`
	ProductID        uuid.UUID  `bun:"product_id,type:uuid"`
	ProductVariantID *uuid.UUID `bun:"product_variant_id,type:uuid"`
	RefSource        *string    `bun:"ref_source"`
	Platform         *string    `bun:"platform"`
	ViewedAt         time.Time  `bun:"viewed_at"`
}
