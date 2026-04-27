package auth

import (
	"net/mail"
	"regexp"

	"go.opentelemetry.io/otel/trace"
)

var (
	registerCitizenIDPattern = regexp.MustCompile(`^\d{13}$`)
	registerPhonePattern     = regexp.MustCompile(`^\+?[0-9]{9,15}$`)
)

type Controller struct {
	tracer trace.Tracer
	svc    *Service
}

func newController(tracer trace.Tracer, svc *Service) *Controller {
	return &Controller{tracer: tracer, svc: svc}
}

type RegisterRequest struct {
	GenderID      string `json:"gender_id"`
	PrefixID      string `json:"prefix_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Displayname   string `json:"displayname"`
	FirstnameTh   string `json:"firstname_th"`
	LastnameTh    string `json:"lastname_th"`
	CitizenID     string `json:"citizen_id"`
	Birthdate     string `json:"birthdate"`
	Phone         string `json:"phone"`
	ProvinceID    string `json:"province_id"`
	DistrictID    string `json:"district_id"`
	SubDistrictID string `json:"sub_district_id"`
	ZipcodeID     string `json:"zipcode_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func isValidEmail(v string) bool {
	if v == "" {
		return false
	}
	addr, err := mail.ParseAddress(v)
	if err != nil {
		return false
	}
	return addr.Address == v
}
