package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShippingMethod struct {
	bun.BaseModel         `bun:"table:shipping_methods,alias:sm"`
	ID                    uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ProviderID            *uuid.UUID `bun:"provider_id,type:uuid"`
	NameTh                *string    `bun:"name_th"`
	NameEn                *string    `bun:"name_en"`
	Code                  *string    `bun:"code"`
	MethodType            *string    `bun:"method_type"`
	BaseFee               float64    `bun:"base_fee,notnull,default:0"`
	FreeShippingThreshold *float64   `bun:"free_shipping_threshold"`
	EstimatedDaysMin      *int       `bun:"estimated_days_min"`
	EstimatedDaysMax      *int       `bun:"estimated_days_max"`
	IsActive              bool       `bun:"is_active,notnull,default:true"`
	CreatedAt             time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt             time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
