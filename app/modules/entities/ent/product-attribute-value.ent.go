package ent

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ProductAttributeValue represents a row in product_attribute_values.
type ProductAttributeValue struct {
	bun.BaseModel `bun:"table:product_attribute_values,alias:pav"`

	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProductID       *uuid.UUID `bun:"product_id,type:uuid"`
	AttributeID     *uuid.UUID `bun:"attribute_id,type:uuid"`
	ValueTh         *string    `bun:"value_th"`
	ValueEn         *string    `bun:"value_en"`
	AdditionalPrice float64    `bun:"additional_price,notnull,default:0"`
}
