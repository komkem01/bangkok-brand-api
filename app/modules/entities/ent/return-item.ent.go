package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReturnItem struct {
	bun.BaseModel     `bun:"table:return_items,alias:ri"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ReturnRequestID   *uuid.UUID `bun:"return_request_id,type:uuid"`
	OrderItemID       *uuid.UUID `bun:"order_item_id,type:uuid"`
	Quantity          int        `bun:"quantity,notnull,default:1"`
	Reason            *string    `bun:"reason"`
	ConditionNote     *string    `bun:"condition_note"`
	EvidenceStorageID *uuid.UUID `bun:"evidence_storage_id,type:uuid"`
	RefundAmount      float64    `bun:"refund_amount,notnull,default:0"`
	CreatedAt         time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
