package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type KYCStatusHistory struct {
	bun.BaseModel     `bun:"table:kyc_status_histories,alias:ksh"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	KYCVerificationID *uuid.UUID `bun:"kyc_verification_id,type:uuid"`
	OldStatus         *string    `bun:"old_status"`
	NewStatus         *string    `bun:"new_status"`
	ChangedByID       *uuid.UUID `bun:"changed_by_id,type:uuid"`
	Remark            *string    `bun:"remark"`
	CreatedAt         time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
