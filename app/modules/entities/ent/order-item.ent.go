package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// OrderItem represents a row in order_items.
type OrderItem struct {
	bun.BaseModel `bun:"table:order_items,alias:oi"`

	ID                  uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrderID             *uuid.UUID     `bun:"order_id,type:uuid"`
	ProductID           *uuid.UUID     `bun:"product_id,type:uuid"`
	ProductNameSnapshot *string        `bun:"product_name_snapshot"`
	SKUSnapshot         *string        `bun:"sku_snapshot"`
	Quantity            *int           `bun:"quantity"`
	UnitPrice           *float64       `bun:"unit_price"`
	SelectedAttributes  map[string]any `bun:"selected_attributes,type:json"`
	SubtotalPrice       *float64       `bun:"subtotal_price"`
	CreatedAt           time.Time      `bun:"created_at,notnull,default:current_timestamp"`
}
