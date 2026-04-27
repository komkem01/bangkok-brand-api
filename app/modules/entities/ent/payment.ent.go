package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Payment represents a row in payments.
type Payment struct {
	bun.BaseModel `bun:"table:payments,alias:p"`

	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrderID           *uuid.UUID `bun:"order_id,type:uuid"`
	PaymentNo         *string    `bun:"payment_no"`
	Method            *string    `bun:"method"`
	Amount            *float64   `bun:"amount"`
	Status            string     `bun:"status,notnull,default:pending"`
	EvidenceStorageID *uuid.UUID `bun:"evidence_storage_id,type:uuid"`
	TransferDateTime  *time.Time `bun:"transfer_date_time"`
	FromBankID        *uuid.UUID `bun:"from_bank_id,type:uuid"`
	TransactionRef    *string    `bun:"transaction_ref"`
	PaidAt            *time.Time `bun:"paid_at"`
	CreatedAt         time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt         time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
