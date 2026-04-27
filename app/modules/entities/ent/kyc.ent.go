package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type KYCVerification struct {
	bun.BaseModel   `bun:"table:kyc_verifications,alias:ky"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MemberID        *uuid.UUID `bun:"member_id,type:uuid"`
	EntityType      string     `bun:"entity_type"`
	LegalName       *string    `bun:"legal_name"`
	BusinessName    *string    `bun:"business_name"`
	CitizenOrTaxID  *string    `bun:"citizen_or_tax_id"`
	ContactPhone    *string    `bun:"contact_phone"`
	ContactEmail    *string    `bun:"contact_email"`
	Status          string     `bun:"status"`
	SubmittedAt     *time.Time `bun:"submitted_at"`
	ReviewedAt      *time.Time `bun:"reviewed_at"`
	ReviewerID      *uuid.UUID `bun:"reviewer_id,type:uuid"`
	RejectionReason *string    `bun:"rejection_reason"`
	CreatedAt       time.Time  `bun:"created_at"`
	UpdatedAt       time.Time  `bun:"updated_at"`
}
