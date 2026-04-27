package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type OrderShipment struct {
	bun.BaseModel      `bun:"table:order_shipments,alias:os"`
	ID                 uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrderID            *uuid.UUID `bun:"order_id,type:uuid"`
	ProviderID         *uuid.UUID `bun:"provider_id,type:uuid"`
	ShippingMethodID   *uuid.UUID `bun:"shipping_method_id,type:uuid"`
	TrackingNumber     *string    `bun:"tracking_number"`
	Status             string     `bun:"status,notnull,default:pending_pickup"`
	ShippingFee        float64    `bun:"shipping_fee,notnull,default:0"`
	ReceiverName       *string    `bun:"receiver_name"`
	ReceiverPhone      *string    `bun:"receiver_phone"`
	ShippingAddress    *string    `bun:"shipping_address"`
	ShippedAt          *time.Time `bun:"shipped_at"`
	ExpectedDeliveryAt *time.Time `bun:"expected_delivery_at"`
	DeliveredAt        *time.Time `bun:"delivered_at"`
	LastStatusAt       *time.Time `bun:"last_status_at"`
	CreatedAt          time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt          time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
