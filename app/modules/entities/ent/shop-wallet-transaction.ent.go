package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ShopWalletTransaction struct {
	bun.BaseModel            `bun:"table:shop_wallet_transactions,alias:swt"`
	ID                       uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ShopID                   *uuid.UUID     `bun:"shop_id,type:uuid"`
	OrderID                  *uuid.UUID     `bun:"order_id,type:uuid"`
	PaymentID                *uuid.UUID     `bun:"payment_id,type:uuid"`
	SettlementBatchID        *uuid.UUID     `bun:"settlement_batch_id,type:uuid"`
	TxType                   *string        `bun:"tx_type"`
	Amount                   float64        `bun:"amount,notnull,default:0"`
	BalanceSnapshot          *float64       `bun:"balance_snapshot"`
	AvailableBalanceSnapshot *float64       `bun:"available_balance_snapshot"`
	Description              *string        `bun:"description"`
	Metadata                 map[string]any `bun:"metadata,type:json"`
	CreatedAt                time.Time      `bun:"created_at,notnull,default:current_timestamp"`
}
