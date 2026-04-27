package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FlashSaleEvent struct {
	bun.BaseModel      `bun:"table:flash_sale_events,alias:fl"`
	ID                 uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID             *uuid.UUID `bun:"shop_id,type:uuid"`
	Name               string     `bun:"name"`
	Description        *string    `bun:"description"`
	CoverStorageID     *uuid.UUID `bun:"cover_storage_id,type:uuid"`
	Status             string     `bun:"status"`
	StartsAt           time.Time  `bun:"starts_at"`
	EndsAt             time.Time  `bun:"ends_at"`
	MaxOrdersPerMember *int       `bun:"max_orders_per_member"`
	IsVisible          bool       `bun:"is_visible"`
	CreatedAt          time.Time  `bun:"created_at"`
	UpdatedAt          time.Time  `bun:"updated_at"`
}
