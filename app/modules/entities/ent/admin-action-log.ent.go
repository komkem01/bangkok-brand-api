package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AdminActionLog struct {
	bun.BaseModel  `bun:"table:admin_action_logs,alias:aal"`
	ID             uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	AdminMemberID  *uuid.UUID     `bun:"admin_member_id,type:uuid"`
	ActionType     *string        `bun:"action_type"`
	ResourceType   *string        `bun:"resource_type"`
	ResourceID     *uuid.UUID     `bun:"resource_id,type:uuid"`
	ShopID         *uuid.UUID     `bun:"shop_id,type:uuid"`
	TargetMemberID *uuid.UUID     `bun:"target_member_id,type:uuid"`
	BeforeData     map[string]any `bun:"before_data,type:json"`
	AfterData      map[string]any `bun:"after_data,type:json"`
	IPAddress      *string        `bun:"ip_address"`
	UserAgent      *string        `bun:"user_agent"`
	CreatedAt      time.Time      `bun:"created_at"`
}
