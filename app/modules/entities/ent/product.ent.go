package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Product represents a row in products.
type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID                 uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CategoryID         *uuid.UUID `bun:"category_id,type:uuid"`
	BrandID            *uuid.UUID `bun:"brand_id,type:uuid"`
	MerchantID         *uuid.UUID `bun:"merchant_id,type:uuid"`
	SKU                *string    `bun:"sku"`
	NameTh             *string    `bun:"name_th"`
	NameEn             *string    `bun:"name_en"`
	ShortDescriptionTh *string    `bun:"short_description_th"`
	FullDescriptionTh  *string    `bun:"full_description_th"`
	Price              *float64   `bun:"price"`
	DiscountPrice      *float64   `bun:"discount_price"`
	IsOnSale           bool       `bun:"is_on_sale,notnull,default:false"`
	Slug               *string    `bun:"slug"`
	MetaTitle          *string    `bun:"meta_title"`
	MetaDescription    *string    `bun:"meta_description"`
	Status             *string    `bun:"status"`
	IsActive           bool       `bun:"is_active,notnull,default:true"`
	IsFeatured         bool       `bun:"is_featured,notnull,default:false"`
	Weight             *float64   `bun:"weight"`
	Width              *float64   `bun:"width"`
	Length             *float64   `bun:"length"`
	Height             *float64   `bun:"height"`
	CreatedAt          time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt          time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt          *time.Time `bun:"deleted_at,soft_delete"`
}
