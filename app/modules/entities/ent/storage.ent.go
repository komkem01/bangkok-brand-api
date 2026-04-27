package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Storage represents a file/media record in the database.
type Storage struct {
	bun.BaseModel `bun:"table:storages,alias:s"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	FileName      *string    `bun:"file_name"`
	FilePath      *string    `bun:"file_path"`
	FileExtension *string    `bun:"file_extension"`
	FileSize      *int       `bun:"file_size"`
	MimeType      *string    `bun:"mime_type"`
	Provider      string     `bun:"provider,notnull,default:'local'"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete"`
}
