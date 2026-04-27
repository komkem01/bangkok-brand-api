package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductVariantValue struct {
	bun.BaseModel `bun:"table:product_variant_values,alias:pvv"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	VariantID     *uuid.UUID `bun:"variant_id,type:uuid"`
	AttributeID   *uuid.UUID `bun:"attribute_id,type:uuid"`
	ValueTh       *string    `bun:"value_th"`
	ValueEn       *string    `bun:"value_en"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
