package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FlashSaleItem struct {
	bun.BaseModel    `bun:"table:flash_sale_items,alias:fsi"`
	ID               uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	FlashSaleEventID *uuid.UUID `bun:"flash_sale_event_id,type:uuid"`
	ProductID        *uuid.UUID `bun:"product_id,type:uuid"`
	ProductVariantID *uuid.UUID `bun:"product_variant_id,type:uuid"`
	OriginalPrice    float64    `bun:"original_price,notnull,default:0"`
	SalePrice        float64    `bun:"sale_price,notnull,default:0"`
	DiscountPercent  *float64   `bun:"discount_percent"`
	StockQuota       *int       `bun:"stock_quota"`
	SoldCount        int        `bun:"sold_count,notnull,default:0"`
	CreatedAt        time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
