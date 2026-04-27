package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChatMessage struct {
	bun.BaseModel `bun:"table:chat_messages,alias:cm"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	RoomID        *uuid.UUID `bun:"room_id,type:uuid"`
	SenderID      *uuid.UUID `bun:"sender_id,type:uuid"`
	Type          string     `bun:"type"`
	Message       *string    `bun:"message"`
	StorageID     *uuid.UUID `bun:"storage_id,type:uuid"`
	IsRead        bool       `bun:"is_read"`
	CreatedAt     time.Time  `bun:"created_at"`
}
