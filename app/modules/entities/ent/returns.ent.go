package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReturnRequest struct {
	bun.BaseModel `bun:"table:return_requests,alias:re"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	RequestNo     *string    `bun:"request_no"`
	OrderID       *uuid.UUID `bun:"order_id,type:uuid"`
	MemberID      *uuid.UUID `bun:"member_id,type:uuid"`
	ShopID        *uuid.UUID `bun:"shop_id,type:uuid"`
	Status        string     `bun:"status"`
	Reason        *string    `bun:"reason"`
	Detail        *string    `bun:"detail"`
	RequestedAt   *time.Time `bun:"requested_at"`
	ApprovedAt    *time.Time `bun:"approved_at"`
	RejectedAt    *time.Time `bun:"rejected_at"`
	ReceivedAt    *time.Time `bun:"received_at"`
	ProcessedByID *uuid.UUID `bun:"processed_by_id,type:uuid"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
}
