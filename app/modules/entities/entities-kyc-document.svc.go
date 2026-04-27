package entities

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.KYCDocumentEntity = (*Service)(nil)

func (s *Service) ListKYCDocuments(ctx context.Context) ([]*ent.KYCDocument, error) {
	var items []*ent.KYCDocument
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetKYCDocumentByID(ctx context.Context, id uuid.UUID) (*ent.KYCDocument, error) {
	item := &ent.KYCDocument{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateKYCDocument(ctx context.Context, item *ent.KYCDocument) (*ent.KYCDocument, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateKYCDocumentByID(ctx context.Context, id uuid.UUID, item *ent.KYCDocument) (*ent.KYCDocument, error) {
	_, err := s.db.NewUpdate().
		Model(item).
		ExcludeColumn("id", "created_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetKYCDocumentByID(ctx, id)
}

func (s *Service) DeleteKYCDocumentByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.KYCDocument)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
