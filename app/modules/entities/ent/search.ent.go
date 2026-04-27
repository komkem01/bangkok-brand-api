package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SearchHistory struct {
	bun.BaseModel `bun:"table:search_histories,alias:se"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	SessionID     *string    `bun:"session_id"`
	Keyword       string     `bun:"keyword"`
	ResultCount   *int       `bun:"result_count"`
	Platform      *string    `bun:"platform"`
	CreatedAt     time.Time  `bun:"created_at"`
}
