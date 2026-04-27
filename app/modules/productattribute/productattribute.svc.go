package productattribute

import (
	entitiesinf "bangkok-brand/app/modules/entities/inf"
	"bangkok-brand/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer trace.Tracer
	db     entitiesinf.ProductAttributeEntity
}

type Config struct{}

type Options struct {
	*config.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.ProductAttributeEntity
}

func newService(opt *Options) *Service {
	return &Service{tracer: opt.tracer, db: opt.db}
}
