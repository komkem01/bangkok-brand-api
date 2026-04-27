package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductReviewImage struct {
	bun.BaseModel `bun:"table:product_review_images,alias:pri"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ReviewID      *uuid.UUID `bun:"review_id,type:uuid"`
	StorageID     *uuid.UUID `bun:"storage_id,type:uuid"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
