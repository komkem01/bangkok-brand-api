package s3

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Config holds MinIO/S3 connection configuration.
type Config struct {
	Endpoint        string `conf:"required"`
	AccessKeyId     string `conf:"required"`
	SecretAccessKey string `conf:"required"`
	BucketName      string
	UseSSL          bool
}

// Service wraps a MinIO client.
type Service struct {
	client     *minio.Client
	bucketName string
}

func newService(conf *Config) *Service {
	client, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyId, conf.SecretAccessKey, ""),
		Secure: conf.UseSSL,
	})
	if err != nil {
		panic(fmt.Sprintf("s3: failed to init minio client: %v", err))
	}
	return &Service{
		client:     client,
		bucketName: conf.BucketName,
	}
}

// PutObject uploads an object to the configured bucket.
func (s *Service) PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	_, err := s.client.PutObject(ctx, s.bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

// PresignedGetObject returns a pre-signed URL valid for the given duration.
func (s *Service) PresignedGetObject(ctx context.Context, objectName string, expiry time.Duration) (*url.URL, error) {
	return s.client.PresignedGetObject(ctx, s.bucketName, objectName, expiry, nil)
}
