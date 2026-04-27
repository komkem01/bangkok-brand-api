package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type LogisticsProvider struct {
	bun.BaseModel       `bun:"table:logistics_providers,alias:lo"`
	ID                  uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	NameTh              *string   `bun:"name_th"`
	NameEn              *string   `bun:"name_en"`
	Code                *string   `bun:"code"`
	TrackingURLTemplate *string   `bun:"tracking_url_template"`
	APIEndpoint         *string   `bun:"api_endpoint"`
	SupportsCOD         bool      `bun:"supports_cod"`
	IsActive            bool      `bun:"is_active"`
	CreatedAt           time.Time `bun:"created_at"`
	UpdatedAt           time.Time `bun:"updated_at"`
}
