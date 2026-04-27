package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Province represents a province record in the database.
type Province struct {
	bun.BaseModel `bun:"table:provinces,alias:prov"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name      string    `bun:"name"`
	IsActive  bool      `bun:"is_active,notnull,default:true"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
