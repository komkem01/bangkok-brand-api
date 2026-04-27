package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type KYCDocument struct {
	bun.BaseModel     `bun:"table:kyc_documents,alias:kd"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	KYCVerificationID *uuid.UUID `bun:"kyc_verification_id,type:uuid"`
	DocumentType      string     `bun:"document_type"`
	StorageID         *uuid.UUID `bun:"storage_id,type:uuid"`
	DocumentNo        *string    `bun:"document_no"`
	IssuedAt          *time.Time `bun:"issued_at"`
	ExpiredAt         *time.Time `bun:"expired_at"`
	IsVerified        bool       `bun:"is_verified,notnull,default:false"`
	VerifiedAt        *time.Time `bun:"verified_at"`
	Note              *string    `bun:"note"`
	CreatedAt         time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
