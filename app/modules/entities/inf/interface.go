package entitiesinf

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"

	"github.com/google/uuid"
)

// ObjectEntity defines the interface for object entity operations such as create, retrieve, update, and soft delete.
type ExampleEntity interface {
	CreateExample(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
	GetExampleByID(ctx context.Context, id uuid.UUID) (*ent.Example, error)
	UpdateExampleByID(ctx context.Context, id uuid.UUID, status ent.ExampleStatus) (*ent.Example, error)
	SoftDeleteExampleByID(ctx context.Context, id uuid.UUID) error
	ListExamplesByStatus(ctx context.Context, status ent.ExampleStatus) ([]*ent.Example, error)
}
type ExampleTwoEntity interface {
	CreateExampleTwo(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
}

// GenderEntity defines CRUD operations for the genders table.
type GenderEntity interface {
	ListGenders(ctx context.Context) ([]*ent.Gender, error)
	GetGenderByID(ctx context.Context, id uuid.UUID) (*ent.Gender, error)
	UpdateGenderByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.Gender, error)
	DeleteGenderByID(ctx context.Context, id uuid.UUID) error
}

// PrefixEntity defines CRUD operations for the prefixes table.
type PrefixEntity interface {
	ListPrefixes(ctx context.Context) ([]*ent.Prefix, error)
	GetPrefixByID(ctx context.Context, id uuid.UUID) (*ent.Prefix, error)
	UpdatePrefixByID(ctx context.Context, id uuid.UUID, genderID *uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.Prefix, error)
	DeletePrefixByID(ctx context.Context, id uuid.UUID) error
}

// ProvinceEntity defines CRUD operations for the provinces table.
type ProvinceEntity interface {
	ListProvinces(ctx context.Context) ([]*ent.Province, error)
	GetProvinceByID(ctx context.Context, id uuid.UUID) (*ent.Province, error)
	UpdateProvinceByID(ctx context.Context, id uuid.UUID, name string, isActive bool) (*ent.Province, error)
	DeleteProvinceByID(ctx context.Context, id uuid.UUID) error
}

// DistrictEntity defines CRUD operations for the districts table.
type DistrictEntity interface {
	ListDistricts(ctx context.Context) ([]*ent.District, error)
	GetDistrictByID(ctx context.Context, id uuid.UUID) (*ent.District, error)
	UpdateDistrictByID(ctx context.Context, id uuid.UUID, provinceID *uuid.UUID, name string, isActive bool) (*ent.District, error)
	DeleteDistrictByID(ctx context.Context, id uuid.UUID) error
}

// SubdistrictEntity defines CRUD operations for the sub_districts table.
type SubdistrictEntity interface {
	ListSubdistricts(ctx context.Context) ([]*ent.Subdistrict, error)
	GetSubdistrictByID(ctx context.Context, id uuid.UUID) (*ent.Subdistrict, error)
	UpdateSubdistrictByID(ctx context.Context, id uuid.UUID, districtID *uuid.UUID, name string, isActive bool) (*ent.Subdistrict, error)
	DeleteSubdistrictByID(ctx context.Context, id uuid.UUID) error
}

// ZipcodeEntity defines CRUD operations for the zipcodes table.
type ZipcodeEntity interface {
	ListZipcodes(ctx context.Context) ([]*ent.Zipcode, error)
	GetZipcodeByID(ctx context.Context, id uuid.UUID) (*ent.Zipcode, error)
	UpdateZipcodeByID(ctx context.Context, id uuid.UUID, subDistrictID *uuid.UUID, name string, isActive bool) (*ent.Zipcode, error)
	DeleteZipcodeByID(ctx context.Context, id uuid.UUID) error
}
