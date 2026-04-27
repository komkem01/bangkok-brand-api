package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChatRoom struct {
	bun.BaseModel `bun:"table:chat_rooms,alias:ch"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrderID       *uuid.UUID `bun:"order_id,type:uuid"`
	BrandID       *uuid.UUID `bun:"brand_id,type:uuid"`
	LastMessage   *string    `bun:"last_message"`
	LastMessageAt *time.Time `bun:"last_message_at"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
}
