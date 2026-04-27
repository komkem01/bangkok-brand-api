package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductVariantStock struct {
	bun.BaseModel     `bun:"table:product_variant_stocks,alias:pvs"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	VariantID         *uuid.UUID `bun:"variant_id,type:uuid"`
	Quantity          int        `bun:"quantity,notnull,default:0"`
	ReservedQuantity  int        `bun:"reserved_quantity,notnull,default:0"`
	LowStockThreshold int        `bun:"low_stock_threshold,notnull,default:5"`
	LastRestockedAt   *time.Time `bun:"last_restocked_at"`
	UpdatedAt         time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
