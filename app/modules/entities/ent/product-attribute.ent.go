package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ProductAttribute represents a row in product_attributes.
type ProductAttribute struct {
	bun.BaseModel `bun:"table:product_attributes,alias:pa"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	NameTh    *string   `bun:"name_th"`
	NameEn    *string   `bun:"name_en"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
