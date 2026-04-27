package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShipmentTrackingHistory struct {
	bun.BaseModel `bun:"table:shipment_tracking_histories,alias:sth"`
	ID            uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShipmentID    *uuid.UUID     `bun:"shipment_id,type:uuid"`
	Status        *string        `bun:"status"`
	Location      *string        `bun:"location"`
	Description   *string        `bun:"description"`
	EventAt       *time.Time     `bun:"event_at"`
	RawPayload    map[string]any `bun:"raw_payload,type:json"`
	CreatedAt     time.Time      `bun:"created_at,notnull,default:current_timestamp"`
}
