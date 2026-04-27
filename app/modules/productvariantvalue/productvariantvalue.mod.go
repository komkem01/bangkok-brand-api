package productvariantvalue

import (
	entitiesinf "bangkok-brand/app/modules/entities/inf"
	"bangkok-brand/internal/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Module struct {
	tracer trace.Tracer
	Svc    *Service
	Ctl    *Controller
}

func New(conf *config.Config[Config], db entitiesinf.ProductVariantValueEntity) *Module {
	tracer := otel.Tracer("bangkok-brand.modules.productvariantvalue")
	svc := newService(&Options{Config: conf, tracer: tracer, db: db})
	return &Module{tracer: tracer, Svc: svc, Ctl: newController(tracer, svc)}
}
