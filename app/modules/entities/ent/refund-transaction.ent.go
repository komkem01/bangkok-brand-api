package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RefundTransaction struct {
	bun.BaseModel   `bun:"table:refund_transactions,alias:rt"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	RefundNo        *string    `bun:"refund_no"`
	ReturnRequestID *uuid.UUID `bun:"return_request_id,type:uuid"`
	OrderID         *uuid.UUID `bun:"order_id,type:uuid"`
	PaymentID       *uuid.UUID `bun:"payment_id,type:uuid"`
	ShopID          *uuid.UUID `bun:"shop_id,type:uuid"`
	Status          string     `bun:"status,notnull,default:pending"`
	Method          *string    `bun:"method"`
	Amount          float64    `bun:"amount,notnull,default:0"`
	GatewayRef      *string    `bun:"gateway_ref"`
	Remark          *string    `bun:"remark"`
	ProcessedByID   *uuid.UUID `bun:"processed_by_id,type:uuid"`
	ProcessedAt     *time.Time `bun:"processed_at"`
	CreatedAt       time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
