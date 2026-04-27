package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// MemberDevice represents a row in member_devices.
type MemberDevice struct {
	bun.BaseModel `bun:"table:member_devices,alias:md"`

	ID          uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID    *uuid.UUID `bun:"member_id,type:uuid"`
	Platform    string     `bun:"platform"`
	DeviceToken string     `bun:"device_token"`
	DeviceName  *string    `bun:"device_name"`
	AppVersion  *string    `bun:"app_version"`
	IsActive    bool       `bun:"is_active,notnull,default:true"`
	LastSeenAt  *time.Time `bun:"last_seen_at"`
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
