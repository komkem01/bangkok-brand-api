package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Brand represents a row in brands.
type Brand struct {
	bun.BaseModel `bun:"table:brands,alias:b"`

	ID          uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	NameTh      *string    `bun:"name_th"`
	NameEn      *string    `bun:"name_en"`
	LogoID      *uuid.UUID `bun:"logo_id,type:uuid"`
	Description *string    `bun:"description"`
	IsActive    bool       `bun:"is_active,notnull,default:true"`
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
