package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ShopMemberEntity = (*Service)(nil)

func (s *Service) ListShopMembers(ctx context.Context) ([]*ent.ShopMember, error) {
	var items []*ent.ShopMember
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetShopMemberByID(ctx context.Context, id uuid.UUID) (*ent.ShopMember, error) {
	item := &ent.ShopMember{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateShopMember(ctx context.Context, item *ent.ShopMember) (*ent.ShopMember, error) {
	_, err := s.db.NewInsert().Model(item).Returning("*").Exec(ctx)
	return item, err
}

func (s *Service) UpdateShopMemberByID(ctx context.Context, id uuid.UUID, item *ent.ShopMember) (*ent.ShopMember, error) {
	item.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(item).ExcludeColumn("id", "created_at").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetShopMemberByID(ctx, id)
}

func (s *Service) DeleteShopMemberByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.ShopMember)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
