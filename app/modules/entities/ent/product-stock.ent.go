package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ProductStock represents a row in product_stocks.
type ProductStock struct {
	bun.BaseModel `bun:"table:product_stocks,alias:ps"`

	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProductID         *uuid.UUID `bun:"product_id,type:uuid"`
	Quantity          int        `bun:"quantity,notnull,default:0"`
	LowStockThreshold int        `bun:"low_stock_threshold,notnull,default:5"`
	LastRestockedAt   *time.Time `bun:"last_restocked_at"`
	UpdatedAt         time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
