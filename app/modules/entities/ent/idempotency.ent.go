package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IdempotencyKey struct {
	bun.BaseModel `bun:"table:idempotency_keys,alias:id"`
	ID            uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Key           string         `bun:"key"`
	Method        string         `bun:"method"`
	Path          string         `bun:"path"`
	RequestHash   string         `bun:"request_hash"`
	StatusCode    *int           `bun:"status_code"`
	ResponseBody  map[string]any `bun:"response_body,type:jsonb"`
	ProcessedAt   *time.Time     `bun:"processed_at"`
	ExpiresAt     time.Time      `bun:"expires_at"`
	CreatedAt     time.Time      `bun:"created_at"`
}
