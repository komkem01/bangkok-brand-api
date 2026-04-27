package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MemberRole string

const (
	MemberRoleCustomer MemberRole = "customer"
	MemberRoleAdmin    MemberRole = "admin"
	MemberRoleMerchant MemberRole = "merchant"
)

type MemberStatus string

const (
	MemberStatusActive    MemberStatus = "active"
	MemberStatusInactive  MemberStatus = "inactive"
	MemberStatusSuspended MemberStatus = "suspended"
)

// Member represents a user account in the members table.
type Member struct {
	bun.BaseModel `bun:"table:members,alias:m"`

	ID             uuid.UUID    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	GenderID       *uuid.UUID   `bun:"gender_id,type:uuid"`
	PrefixID       *uuid.UUID   `bun:"prefix_id,type:uuid"`
	Email          *string      `bun:"email"`
	Password       *string      `bun:"password"`
	MemberNo       *string      `bun:"member_no"`
	ProfileImageID *uuid.UUID   `bun:"profile_image_id,type:uuid"`
	Displayname    *string      `bun:"displayname"`
	FirstnameTh    *string      `bun:"firstname_th"`
	LastnameTh     *string      `bun:"lastname_th"`
	CitizenID      *string      `bun:"citizen_id"`
	Birthdate      *time.Time   `bun:"birthdate,type:date"`
	Phone          *string      `bun:"phone"`
	Role           MemberRole   `bun:"role,notnull,default:'customer'"`
	Status         MemberStatus `bun:"status,notnull,default:'active'"`
	ProvinceID     *uuid.UUID   `bun:"province_id,type:uuid"`
	DistrictID     *uuid.UUID   `bun:"district_id,type:uuid"`
	SubDistrictID  *uuid.UUID   `bun:"sub_district_id,type:uuid"`
	ZipcodeID      *uuid.UUID   `bun:"zipcode_id,type:uuid"`
	RegisterdAt    *time.Time   `bun:"registerd_at"`
	LastedLogin    *time.Time   `bun:"lasted_login"`
	IsVerified     bool         `bun:"is_verified,notnull,default:false"`
	TotalPoints    int          `bun:"total_points,notnull,default:0"`
	CreatedAt      time.Time    `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time    `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt      *time.Time   `bun:"deleted_at,soft_delete"`
}
