package payment

import (
	"context"
	"database/sql"
	"errors"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	tracer trace.Tracer
	svc    *Service
}

type idRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func newController(tracer trace.Tracer, svc *Service) *Controller {
	return &Controller{tracer: tracer, svc: svc}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("payment.list.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, items)
}

func (c *Controller) Info(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item, err := c.svc.Info(ctx.Request.Context(), uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		log.Errf("payment.info.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, item)
}

func (c *Controller) Create(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	var item ent.Payment
	if err := ctx.ShouldBindJSON(&item); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	created, err := c.svc.Create(ctx.Request.Context(), &item)
	if err != nil {
		log.Errf("payment.create.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, created)
}

func (c *Controller) Update(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	var item ent.Payment
	if err := ctx.ShouldBindJSON(&item); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	updated, err := c.svc.Update(ctx.Request.Context(), uuid.MustParse(req.ID), &item)
	if err != nil {
		log.Errf("payment.update.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, updated)
}

func (c *Controller) Delete(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	err := c.svc.Delete(ctx.Request.Context(), uuid.MustParse(req.ID))
	if err != nil {
		log.Errf("payment.delete.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, nil)
}

func (s *Service) List(ctx context.Context) ([]*ent.Payment, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "payment.List")
	defer span.End()
	items, err := s.db.ListPayments(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("payment.list.error: %v", err)
		return nil, err
	}
	return items, nil
}

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Payment, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "payment.Info")
	defer span.End()
	item, err := s.db.GetPaymentByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("payment.info.error: %v", err)
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(ctx context.Context, item *ent.Payment) (*ent.Payment, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "payment.Create")
	defer span.End()
	created, err := s.db.CreatePayment(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("payment.create.error: %v", err)
		return nil, err
	}
	return created, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.Payment) (*ent.Payment, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "payment.Update")
	defer span.End()
	updated, err := s.db.UpdatePaymentByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("payment.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "payment.Delete")
	defer span.End()
	err := s.db.DeletePaymentByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("payment.delete.error: %v", err)
		return err
	}
	return nil
}
