package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Wishlist struct {
	bun.BaseModel    `bun:"table:wishlists,alias:wi"`
	ID               uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID         uuid.UUID  `bun:"member_id,type:uuid"`
	ProductID        uuid.UUID  `bun:"product_id,type:uuid"`
	ProductVariantID *uuid.UUID `bun:"product_variant_id,type:uuid"`
	CreatedAt        time.Time  `bun:"created_at"`
}
