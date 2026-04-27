package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChatParticipant struct {
	bun.BaseModel `bun:"table:chat_participants,alias:cp"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	RoomID        *uuid.UUID `bun:"room_id,type:uuid"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	JoinedAt      *time.Time `bun:"joined_at"`
	LastReadAt    *time.Time `bun:"last_read_at"`
	IsActive      bool       `bun:"is_active"`
}
