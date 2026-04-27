package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShopMember struct {
	bun.BaseModel `bun:"table:shop_members,alias:sm"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID        *uuid.UUID `bun:"shop_id,type:uuid"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	Role          string     `bun:"role,notnull,default:staff"`
	IsActive      bool       `bun:"is_active,notnull,default:true"`
	JoinedAt      *time.Time `bun:"joined_at"`
	InvitedByID   *uuid.UUID `bun:"invited_by_id,type:uuid"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
