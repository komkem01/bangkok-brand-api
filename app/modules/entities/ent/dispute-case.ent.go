package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DisputeCase struct {
	bun.BaseModel   `bun:"table:dispute_cases,alias:dc"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CaseNo          *string    `bun:"case_no"`
	OrderID         *uuid.UUID `bun:"order_id,type:uuid"`
	ReturnRequestID *uuid.UUID `bun:"return_request_id,type:uuid"`
	MemberID        *uuid.UUID `bun:"member_id,type:uuid"`
	ShopID          *uuid.UUID `bun:"shop_id,type:uuid"`
	Subject         *string    `bun:"subject"`
	Detail          *string    `bun:"detail"`
	Status          string     `bun:"status"`
	AssignedAdminID *uuid.UUID `bun:"assigned_admin_id,type:uuid"`
	OpenedAt        *time.Time `bun:"opened_at"`
	ClosedAt        *time.Time `bun:"closed_at"`
	ResolutionNote  *string    `bun:"resolution_note"`
	CreatedAt       time.Time  `bun:"created_at"`
	UpdatedAt       time.Time  `bun:"updated_at"`
}
