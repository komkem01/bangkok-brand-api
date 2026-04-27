package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ContactEntity = (*Service)(nil)

func (s *Service) ListContacts(ctx context.Context) ([]*ent.MemberContact, error) {
	var contacts []*ent.MemberContact
	err := s.db.NewSelect().
		Model(&contacts).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return contacts, err
}

func (s *Service) GetContactByID(ctx context.Context, id uuid.UUID) (*ent.MemberContact, error) {
	contact := &ent.MemberContact{}
	err := s.db.NewSelect().
		Model(contact).
		Where("id = ?", id).
		Scan(ctx)
	return contact, err
}

func (s *Service) CreateContact(ctx context.Context, c *ent.MemberContact) (*ent.MemberContact, error) {
	_, err := s.db.NewInsert().
		Model(c).
		Returning("*").
		Exec(ctx)
	return c, err
}

func (s *Service) UpdateContactByID(ctx context.Context, id uuid.UUID, c *ent.MemberContact) (*ent.MemberContact, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.MemberContact)(nil)).
		Set("member_id = ?", c.MemberID).
		Set("contact_type_id = ?", c.ContactTypeID).
		Set("value = ?", c.Value).
		Set("is_primary = ?", c.IsPrimary).
		Set("is_verified = ?", c.IsVerified).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetContactByID(ctx, id)
}

func (s *Service) DeleteContactByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.MemberContact)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
