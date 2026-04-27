package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// District represents a district (อำเภอ/เขต) record in the database.
type District struct {
	bun.BaseModel `bun:"table:districts,alias:dist"`

	ID         uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProvinceID *uuid.UUID `bun:"province_id,type:uuid"`
	Name       string     `bun:"name"`
	IsActive   bool       `bun:"is_active,notnull,default:true"`
	CreatedAt  time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt  time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
