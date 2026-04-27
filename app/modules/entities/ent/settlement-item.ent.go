package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SettlementItem struct {
	bun.BaseModel `bun:"table:settlement_items,alias:si"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	BatchID       *uuid.UUID `bun:"batch_id,type:uuid"`
	ShopID        *uuid.UUID `bun:"shop_id,type:uuid"`
	OrderID       *uuid.UUID `bun:"order_id,type:uuid"`
	PaymentID     *uuid.UUID `bun:"payment_id,type:uuid"`
	Status        string     `bun:"status,notnull,default:pending"`
	GrossAmount   float64    `bun:"gross_amount,notnull,default:0"`
	FeeAmount     float64    `bun:"fee_amount,notnull,default:0"`
	NetAmount     float64    `bun:"net_amount,notnull,default:0"`
	Note          *string    `bun:"note"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
