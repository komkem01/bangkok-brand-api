package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// MemberBankAccount represents a row in member_bank_accounts.
type MemberBankAccount struct {
	bun.BaseModel `bun:"table:member_bank_accounts,alias:mba"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	BankID        *uuid.UUID `bun:"bank_id,type:uuid"`
	AccountNumber *string    `bun:"account_number"`
	AccountName   *string    `bun:"account_name"`
	BranchName    *string    `bun:"branch_name"`
	IsDefault     bool       `bun:"is_default,notnull,default:false"`
	IsVerified    bool       `bun:"is_verified,notnull,default:false"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}
