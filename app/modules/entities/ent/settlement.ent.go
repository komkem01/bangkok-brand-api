package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SettlementBatch struct {
	bun.BaseModel    `bun:"table:settlement_batches,alias:se"`
	ID               uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID           *uuid.UUID `bun:"shop_id,type:uuid"`
	Status           string     `bun:"status"`
	PeriodStart      *time.Time `bun:"period_start"`
	PeriodEnd        *time.Time `bun:"period_end"`
	GrossAmount      float64    `bun:"gross_amount"`
	FeeAmount        float64    `bun:"fee_amount"`
	AdjustmentAmount float64    `bun:"adjustment_amount"`
	NetAmount        float64    `bun:"net_amount"`
	TransferFee      float64    `bun:"transfer_fee"`
	FinalAmount      float64    `bun:"final_amount"`
	PayoutAccountID  *uuid.UUID `bun:"payout_account_id,type:uuid"`
	RequestedByID    *uuid.UUID `bun:"requested_by_id,type:uuid"`
	ApprovedByID     *uuid.UUID `bun:"approved_by_id,type:uuid"`
	PaidByID         *uuid.UUID `bun:"paid_by_id,type:uuid"`
	PayoutReference  *string    `bun:"payout_reference"`
	PaidAt           *time.Time `bun:"paid_at"`
	Note             *string    `bun:"note"`
	CreatedAt        time.Time  `bun:"created_at"`
	UpdatedAt        time.Time  `bun:"updated_at"`
}
