package pointsetting

import (
	entitiesinf "bangkok-brand/app/modules/entities/inf"
	"bangkok-brand/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Config struct{}

type Options struct {
	Config *config.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.PointSettingEntity
}

type Service struct {
	*Options
}

func newService(opts *Options) *Service {
	return &Service{Options: opts}
}
