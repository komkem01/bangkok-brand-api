package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AddressType represents a row in address_types.
type AddressType struct {
	bun.BaseModel `bun:"table:address_types,alias:at"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	NameTh    string    `bun:"name_th"`
	NameEn    string    `bun:"name_en"`
	IsActive  bool      `bun:"is_active,notnull,default:true"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
