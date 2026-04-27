package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// CartItem represents a row in cart_items.
type CartItem struct {
	bun.BaseModel `bun:"table:cart_items,alias:ci"`

	ID                      uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CartID                  *uuid.UUID     `bun:"cart_id,type:uuid"`
	ProductID               *uuid.UUID     `bun:"product_id,type:uuid"`
	Quantity                int            `bun:"quantity,notnull,default:1"`
	SelectedAttributeValues map[string]any `bun:"selected_attribute_values,type:json"`
	UnitPrice               *float64       `bun:"unit_price"`
	SubtotalPrice           *float64       `bun:"subtotal_price"`
	CreatedAt               time.Time      `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt               time.Time      `bun:"updated_at,notnull,default:current_timestamp"`
}
