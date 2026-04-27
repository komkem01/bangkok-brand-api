package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type WebhookEvent struct {
	bun.BaseModel    `bun:"table:webhook_events,alias:we"`
	ID               uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	EventType        string         `bun:"event_type"`
	Payload          map[string]any `bun:"payload,type:jsonb"`
	Status           string         `bun:"status"`
	ShopID           *uuid.UUID     `bun:"shop_id,type:uuid"`
	EndpointURL      string         `bun:"endpoint_url"`
	SecretHash       *string        `bun:"secret_hash"`
	AttemptCount     int            `bun:"attempt_count"`
	MaxAttempts      int            `bun:"max_attempts"`
	LastAttemptAt    *time.Time     `bun:"last_attempt_at"`
	NextRetryAt      *time.Time     `bun:"next_retry_at"`
	LastResponseCode *int           `bun:"last_response_code"`
	LastResponseBody *string        `bun:"last_response_body"`
	DeliveredAt      *time.Time     `bun:"delivered_at"`
	CreatedAt        time.Time      `bun:"created_at"`
	UpdatedAt        time.Time      `bun:"updated_at"`
}
