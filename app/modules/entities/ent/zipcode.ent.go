package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Zipcode represents a postal code (รหัสไปรษณีย์) record in the database.
type Zipcode struct {
	bun.BaseModel `bun:"table:zipcodes,alias:zip"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SubDistrictID *uuid.UUID `bun:"sub_district_id,type:uuid"`
	Name          string     `bun:"name"`
	IsActive      bool       `bun:"is_active,notnull,default:true"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
