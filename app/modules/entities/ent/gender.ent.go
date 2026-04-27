package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Gender represents a gender record in the database.
type Gender struct {
	bun.BaseModel `bun:"table:genders,alias:g"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	NameTh    string     `bun:"name_th"`
	NameEn    string     `bun:"name_en"`
	IsActive  bool       `bun:"is_active,notnull,default:true"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
}
