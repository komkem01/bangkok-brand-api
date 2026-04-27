package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Invoice struct {
	bun.BaseModel  `bun:"table:invoices,alias:in"`
	ID             uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	InvoiceNo      *string    `bun:"invoice_no"`
	OrderID        *uuid.UUID `bun:"order_id,type:uuid"`
	PaymentID      *uuid.UUID `bun:"payment_id,type:uuid"`
	MemberID       *uuid.UUID `bun:"member_id,type:uuid"`
	ShopID         *uuid.UUID `bun:"shop_id,type:uuid"`
	Type           string     `bun:"type"`
	Status         string     `bun:"status"`
	IssueDate      *time.Time `bun:"issue_date,type:date"`
	DueDate        *time.Time `bun:"due_date,type:date"`
	SubTotal       float64    `bun:"sub_total"`
	DiscountAmount float64    `bun:"discount_amount"`
	TaxAmount      float64    `bun:"tax_amount"`
	TotalAmount    float64    `bun:"total_amount"`
	TaxRate        float64    `bun:"tax_rate"`
	BillingName    *string    `bun:"billing_name"`
	BillingAddress *string    `bun:"billing_address"`
	BillingTaxID   *string    `bun:"billing_tax_id"`
	Note           *string    `bun:"note"`
	CreatedAt      time.Time  `bun:"created_at"`
	UpdatedAt      time.Time  `bun:"updated_at"`
}
