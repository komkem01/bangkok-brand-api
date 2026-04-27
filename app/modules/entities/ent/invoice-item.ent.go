package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type InvoiceItem struct {
	bun.BaseModel  `bun:"table:invoice_items,alias:ii"`
	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	InvoiceID      *uuid.UUID `bun:"invoice_id,type:uuid"`
	OrderItemID    *uuid.UUID `bun:"order_item_id,type:uuid"`
	Description    *string    `bun:"description"`
	Quantity       int        `bun:"quantity,notnull,default:1"`
	UnitPrice      float64    `bun:"unit_price,notnull,default:0"`
	DiscountAmount float64    `bun:"discount_amount,notnull,default:0"`
	TaxAmount      float64    `bun:"tax_amount,notnull,default:0"`
	TotalAmount    float64    `bun:"total_amount,notnull,default:0"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
