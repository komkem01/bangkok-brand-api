package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Category represents a row in categories.
type Category struct {
	bun.BaseModel `bun:"table:categories,alias:c"`

	ID          uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ParentID    *uuid.UUID `bun:"parent_id,type:uuid"`
	NameTh      *string    `bun:"name_th"`
	NameEn      *string    `bun:"name_en"`
	Description *string    `bun:"description"`
	ImageID     *uuid.UUID `bun:"image_id,type:uuid"`
	Slug        *string    `bun:"slug"`
	IsActive    bool       `bun:"is_active,notnull,default:true"`
	SortOrder   int        `bun:"sort_order,notnull,default:0"`
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
