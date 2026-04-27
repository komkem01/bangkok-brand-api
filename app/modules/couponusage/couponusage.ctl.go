package couponusage

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
		log.Errf("couponusage.list.error: %v", err)
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
		log.Errf("couponusage.info.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, item)
}

func (c *Controller) Create(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	var item ent.CouponUsage
	if err := ctx.ShouldBindJSON(&item); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	created, err := c.svc.Create(ctx.Request.Context(), &item)
	if err != nil {
		log.Errf("couponusage.create.error: %v", err)
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
	var item ent.CouponUsage
	if err := ctx.ShouldBindJSON(&item); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	updated, err := c.svc.Update(ctx.Request.Context(), uuid.MustParse(req.ID), &item)
	if err != nil {
		log.Errf("couponusage.update.error: %v", err)
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
		log.Errf("couponusage.delete.error: %v", err)
		base.InternalServerError(ctx, i18n.InternalServerError, nil)
		return
	}
	base.Success(ctx, nil)
}

func (s *Service) List(ctx context.Context) ([]*ent.CouponUsage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "couponusage.List")
	defer span.End()
	items, err := s.db.ListCouponUsages(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("couponusage.list.error: %v", err)
		return nil, err
	}
	return items, nil
}

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.CouponUsage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "couponusage.Info")
	defer span.End()
	item, err := s.db.GetCouponUsageByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("couponusage.info.error: %v", err)
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(ctx context.Context, item *ent.CouponUsage) (*ent.CouponUsage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "couponusage.Create")
	defer span.End()
	created, err := s.db.CreateCouponUsage(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("couponusage.create.error: %v", err)
		return nil, err
	}
	return created, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.CouponUsage) (*ent.CouponUsage, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "couponusage.Update")
	defer span.End()
	updated, err := s.db.UpdateCouponUsageByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("couponusage.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "couponusage.Delete")
	defer span.End()
	err := s.db.DeleteCouponUsageByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("couponusage.delete.error: %v", err)
		return err
	}
	return nil
}
