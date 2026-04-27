package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.StorageEntity = (*Service)(nil)

// ListStorages returns all non-deleted storages ordered by created_at DESC.
func (s *Service) ListStorages(ctx context.Context) ([]*ent.Storage, error) {
	var storages []*ent.Storage
	err := s.db.NewSelect().
		Model(&storages).
		Where("deleted_at IS NULL").
		OrderExpr("created_at DESC").
		Scan(ctx)
	return storages, err
}

// GetStorageByID returns a single storage by its primary key.
func (s *Service) GetStorageByID(ctx context.Context, id uuid.UUID) (*ent.Storage, error) {
	storage := &ent.Storage{}
	err := s.db.NewSelect().
		Model(storage).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	return storage, err
}

// DeleteStorageByID soft-deletes a storage record.
func (s *Service) DeleteStorageByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewUpdate().
		Model((*ent.Storage)(nil)).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}

// CreateStorage inserts a new storage record and returns it with the generated ID.
func (s *Service) CreateStorage(ctx context.Context, st *ent.Storage) (*ent.Storage, error) {
	_, err := s.db.NewInsert().
		Model(st).
		Returning("*").
		Exec(ctx)
	return st, err
}
