package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShippingZone struct {
	bun.BaseModel `bun:"table:shipping_zones,alias:sh"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID        *uuid.UUID `bun:"shop_id,type:uuid"`
	NameTh        *string    `bun:"name_th"`
	NameEn        *string    `bun:"name_en"`
	Description   *string    `bun:"description"`
	IsActive      bool       `bun:"is_active"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
}
