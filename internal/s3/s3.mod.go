package s3

// Module is the S3/MinIO module.
type Module struct {
	Svc *Service
}

// New initialises the MinIO client from the given config.
func New(conf *Config) *Module {
	svc := newService(conf)
	return &Module{Svc: svc}
}
