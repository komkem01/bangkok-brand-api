package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// OrderStatusHistory represents a row in order_status_histories.
type OrderStatusHistory struct {
	bun.BaseModel `bun:"table:order_status_histories,alias:osh"`

	ID          uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OrderID     *uuid.UUID `bun:"order_id,type:uuid"`
	Status      *string    `bun:"status"`
	Remark      *string    `bun:"remark"`
	ChangedByID *uuid.UUID `bun:"changed_by_id,type:uuid"`
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
}
