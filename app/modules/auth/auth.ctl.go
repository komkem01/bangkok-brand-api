package auth

import (
	"errors"
	"net/mail"
	"strings"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	tracer trace.Tracer
	svc    *Service
}

func newController(tracer trace.Tracer, svc *Service) *Controller {
	return &Controller{tracer: tracer, svc: svc}
}

type RegisterRequest struct {
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Phone       *string `json:"phone"`
	FirstnameTh *string `json:"firstname_th"`
	LastnameTh  *string `json:"lastname_th"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (c *Controller) GoogleLogin(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)
	defer span.End()
	base.NotImplemented(ctx, i18n.BadRequest, nil)
}

func (c *Controller) GoogleCallback(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)
	defer span.End()
	base.NotImplemented(ctx, i18n.BadRequest, nil)
}

func (c *Controller) Register(ctx *gin.Context) {
	var body RegisterRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if !isValidEmail(strings.TrimSpace(body.Email)) {
		base.BadRequest(ctx, i18n.MemberInvalidEmail, nil)
		return
	}

	result, err := c.svc.Register(ctx.Request.Context(), &RegisterInput{
		Email:     body.Email,
		Password:  body.Password,
		Phone:     body.Phone,
		Firstname: body.FirstnameTh,
		Lastname:  body.LastnameTh,
	})
	if err != nil {
		switch {
		case errors.Is(err, ErrEmailExists):
			base.BadRequest(ctx, i18n.BadRequest, nil)
		case errors.Is(err, ErrInvalidCredentials):
			base.BadRequest(ctx, i18n.BadRequest, nil)
		default:
			log.Errf("auth.register.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, result)
}

func (c *Controller) Login(ctx *gin.Context) {
	var body LoginRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if !isValidEmail(strings.TrimSpace(body.Email)) {
		base.BadRequest(ctx, i18n.MemberInvalidEmail, nil)
		return
	}

	result, err := c.svc.Login(ctx.Request.Context(), &LoginInput{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidCredentials):
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		default:
			log.Errf("auth.login.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, result)
}

func (c *Controller) RefreshToken(ctx *gin.Context) {
	var body RefreshTokenRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	body.RefreshToken = strings.TrimSpace(body.RefreshToken)
	if body.RefreshToken == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	tokens, err := c.svc.RefreshToken(ctx.Request.Context(), body.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidToken):
			base.Unauthorized(ctx, i18n.Unauthorized, nil)
		case errors.Is(err, ErrMemberNotActive):
			base.Forbidden(ctx, i18n.Forbidden, nil)
		default:
			log.Errf("auth.refresh.error: %v", err)
			base.InternalServerError(ctx, i18n.InternalServerError, nil)
		}
		return
	}

	base.Success(ctx, tokens)
}

func (c *Controller) Logout(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := c.svc.Logout(ctx.Request.Context()); err != nil {
		log.Errf("auth.logout.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}

	base.Success(ctx, nil)
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
