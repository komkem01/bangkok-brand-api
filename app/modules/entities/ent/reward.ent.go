package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Reward struct {
	bun.BaseModel  `bun:"table:rewards,alias:rw"`
	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	NameTh         *string    `bun:"name_th"`
	NameEn         *string    `bun:"name_en"`
	Description    *string    `bun:"description"`
	Type           *string    `bun:"type"`
	PointsRequired *int       `bun:"points_required"`
	ImageID        *uuid.UUID `bun:"image_id,type:uuid"`
	StockQuantity  int        `bun:"stock_quantity,notnull,default:0"`
	StartDate      *time.Time `bun:"start_date"`
	EndDate        *time.Time `bun:"end_date"`
	IsActive       bool       `bun:"is_active,notnull,default:true"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
