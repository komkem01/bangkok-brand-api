package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductVariant struct {
	bun.BaseModel   `bun:"table:product_variants,alias:va"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProductID       *uuid.UUID `bun:"product_id,type:uuid"`
	ShopID          *uuid.UUID `bun:"shop_id,type:uuid"`
	SKU             *string    `bun:"sku"`
	NameTh          *string    `bun:"name_th"`
	NameEn          *string    `bun:"name_en"`
	Barcode         *string    `bun:"barcode"`
	Price           *float64   `bun:"price"`
	DiscountPrice   *float64   `bun:"discount_price"`
	AdditionalPrice float64    `bun:"additional_price"`
	IsDefault       bool       `bun:"is_default"`
	IsActive        bool       `bun:"is_active"`
	Weight          *float64   `bun:"weight"`
	Width           *float64   `bun:"width"`
	Length          *float64   `bun:"length"`
	Height          *float64   `bun:"height"`
	CreatedAt       time.Time  `bun:"created_at"`
	UpdatedAt       time.Time  `bun:"updated_at"`
}
