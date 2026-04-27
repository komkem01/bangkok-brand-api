package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Shop struct {
	bun.BaseModel     `bun:"table:shops,alias:sh"`
	ID                uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	OwnerMemberID     *uuid.UUID `bun:"owner_member_id,type:uuid"`
	BrandID           *uuid.UUID `bun:"brand_id,type:uuid"`
	KYCVerificationID *uuid.UUID `bun:"kyc_verification_id,type:uuid"`
	ShopCode          *string    `bun:"shop_code"`
	NameTh            *string    `bun:"name_th"`
	NameEn            *string    `bun:"name_en"`
	Slug              *string    `bun:"slug"`
	Description       *string    `bun:"description"`
	LogoID            *uuid.UUID `bun:"logo_id,type:uuid"`
	CoverImageID      *uuid.UUID `bun:"cover_image_id,type:uuid"`
	ContactPhone      *string    `bun:"contact_phone"`
	ContactEmail      *string    `bun:"contact_email"`
	AddressDetail     *string    `bun:"address_detail"`
	ProvinceID        *uuid.UUID `bun:"province_id,type:uuid"`
	DistrictID        *uuid.UUID `bun:"district_id,type:uuid"`
	SubDistrictID     *uuid.UUID `bun:"sub_district_id,type:uuid"`
	ZipcodeID         *uuid.UUID `bun:"zipcode_id,type:uuid"`
	Status            string     `bun:"status"`
	IsActive          bool       `bun:"is_active"`
	OpenedAt          *time.Time `bun:"opened_at"`
	CreatedAt         time.Time  `bun:"created_at"`
	UpdatedAt         time.Time  `bun:"updated_at"`
}
