package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ProductImage represents a row in product_images.
type ProductImage struct {
	bun.BaseModel `bun:"table:product_images,alias:pi"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProductID *uuid.UUID `bun:"product_id,type:uuid"`
	StorageID *uuid.UUID `bun:"storage_id,type:uuid"`
	IsMain    bool       `bun:"is_main,notnull,default:false"`
	SortOrder int        `bun:"sort_order,notnull,default:0"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
