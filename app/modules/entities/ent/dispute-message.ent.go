package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DisputeMessage struct {
	bun.BaseModel     `bun:"table:dispute_messages,alias:dm"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	DisputeCaseID     *uuid.UUID `bun:"dispute_case_id,type:uuid"`
	SenderMemberID    *uuid.UUID `bun:"sender_member_id,type:uuid"`
	Party             string     `bun:"party"`
	Message           *string    `bun:"message"`
	EvidenceStorageID *uuid.UUID `bun:"evidence_storage_id,type:uuid"`
	IsInternalNote    bool       `bun:"is_internal_note,notnull,default:false"`
	CreatedAt         time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
