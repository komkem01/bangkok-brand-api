package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Notification struct {
	bun.BaseModel `bun:"table:notifications,alias:no"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	Channel       string     `bun:"channel"`
	Status        string     `bun:"status"`
	Title         string     `bun:"title"`
	Body          *string    `bun:"body"`
	ImageURL      *string    `bun:"image_url"`
	ActionURL     *string    `bun:"action_url"`
	RefType       *string    `bun:"ref_type"`
	RefID         *uuid.UUID `bun:"ref_id,type:uuid"`
	ReadAt        *time.Time `bun:"read_at"`
	SentAt        *time.Time `bun:"sent_at"`
	FailedReason  *string    `bun:"failed_reason"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
}
