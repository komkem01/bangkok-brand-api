package wishlist

import "go.opentelemetry.io/otel/trace"

type Controller struct {
	tracer trace.Tracer
	svc    *Service
}

func newController(tracer trace.Tracer, svc *Service) *Controller {
	return &Controller{tracer: tracer, svc: svc}
}

type idRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
