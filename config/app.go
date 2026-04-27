package config

import (
	"bangkok-brand/app/modules/example"
	exampletwo "bangkok-brand/app/modules/example-two"
	"bangkok-brand/app/modules/sentry"
	"bangkok-brand/app/modules/specs"
	"bangkok-brand/internal/kafka"
	"bangkok-brand/internal/log"
	"bangkok-brand/internal/otel/collector"
)

// Config is a struct that contains all the configuration of the application.
type Config struct {
	Database Database

	AppName     string
	AppKey      string
	Environment string
	Specs       specs.Config
	Debug       bool

	Port           int
	HttpJsonNaming string

	SslCaPath      string
	SslPrivatePath string
	SslCertPath    string

	Otel   collector.Config
	Sentry sentry.Config

	Kafka kafka.Config
	Log   log.Option

	Example example.Config

	ExampleTwo exampletwo.Config
}

var App = Config{
	Specs: specs.Config{
		Version: "v1",
	},
	Database: database,
	Kafka:    kafkaConf,

	AppName: "go_app",
	Port:    8080,
	AppKey:  "secret",
	Debug:   false,

	HttpJsonNaming: "snake_case",

	SslCaPath:      "bangkok-brand/cert/ca.pem",
	SslPrivatePath: "bangkok-brand/cert/server.pem",
	SslCertPath:    "bangkok-brand/cert/server-key.pem",

	Otel: collector.Config{
		CollectorEndpoint: "",
		LogMode:           "noop",
		TraceMode:         "noop",
		MetricMode:        "noop",
		TraceRatio:        0.01,
	},
}
