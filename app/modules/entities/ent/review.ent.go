package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductReview struct {
	bun.BaseModel      `bun:"table:product_reviews,alias:re"`
	ID                 uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProductID          *uuid.UUID `bun:"product_id,type:uuid"`
	MemberID           *uuid.UUID `bun:"member_id,type:uuid"`
	OrderID            *uuid.UUID `bun:"order_id,type:uuid"`
	Rating             *int       `bun:"rating"`
	Comment            *string    `bun:"comment"`
	IsVerifiedPurchase bool       `bun:"is_verified_purchase"`
	IsHidden           bool       `bun:"is_hidden"`
	CreatedAt          time.Time  `bun:"created_at"`
	UpdatedAt          time.Time  `bun:"updated_at"`
}
