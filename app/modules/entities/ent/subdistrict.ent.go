package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Subdistrict represents a sub-district (ตำบล/แขวง) record in the database.
type Subdistrict struct {
	bun.BaseModel `bun:"table:sub_districts,alias:subdist"`

	ID         uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	DistrictID *uuid.UUID `bun:"district_id,type:uuid"`
	Name       string     `bun:"name"`
	IsActive   bool       `bun:"is_active,notnull,default:true"`
	CreatedAt  time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt  time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
