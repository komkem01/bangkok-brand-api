package subdistrict

import (
	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	tracer trace.Tracer
	svc    *Service
}

func newController(tracer trace.Tracer, svc *Service) *Controller {
	return &Controller{
		tracer: tracer,
		svc:    svc,
	}
}
