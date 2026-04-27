package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// MemberContact represents a row in member_contacts.
type MemberContact struct {
	bun.BaseModel `bun:"table:member_contacts,alias:mc"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	ContactTypeID *uuid.UUID `bun:"contact_type_id,type:uuid"`
	Value         *string    `bun:"value"`
	IsPrimary     bool       `bun:"is_primary,notnull,default:false"`
	IsVerified    bool       `bun:"is_verified,notnull,default:false"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
