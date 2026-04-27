package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AuditLog struct {
	bun.BaseModel `bun:"table:audit_logs,alias:au"`
	ID            uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	TableName     string         `bun:"table_name"`
	RecordID      *uuid.UUID     `bun:"record_id,type:uuid"`
	Action        string         `bun:"action"`
	ActorID       *uuid.UUID     `bun:"actor_id,type:uuid"`
	ActorType     *string        `bun:"actor_type"`
	OldValues     map[string]any `bun:"old_values,type:jsonb"`
	NewValues     map[string]any `bun:"new_values,type:jsonb"`
	ChangedFields []string       `bun:"changed_fields,array"`
	IPAddress     *string        `bun:"ip_address"`
	UserAgent     *string        `bun:"user_agent"`
	RequestID     *string        `bun:"request_id"`
	CreatedAt     time.Time      `bun:"created_at"`
}
