package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.MemberDeviceEntity = (*Service)(nil)

func (s *Service) ListMemberDevices(ctx context.Context) ([]*ent.MemberDevice, error) {
	var items []*ent.MemberDevice
	err := s.db.NewSelect().
		Model(&items).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) ListMemberDevicesByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberDevice, error) {
	var items []*ent.MemberDevice
	err := s.db.NewSelect().
		Model(&items).
		Where("member_id = ?", memberID).
		OrderExpr("created_at DESC").
		Scan(ctx)
	return items, err
}

func (s *Service) GetMemberDeviceByID(ctx context.Context, id uuid.UUID) (*ent.MemberDevice, error) {
	item := &ent.MemberDevice{}
	err := s.db.NewSelect().
		Model(item).
		Where("id = ?", id).
		Scan(ctx)
	return item, err
}

func (s *Service) CreateMemberDevice(ctx context.Context, item *ent.MemberDevice) (*ent.MemberDevice, error) {
	_, err := s.db.NewInsert().
		Model(item).
		Returning("*").
		Exec(ctx)
	return item, err
}

func (s *Service) UpdateMemberDeviceByID(ctx context.Context, id uuid.UUID, item *ent.MemberDevice) (*ent.MemberDevice, error) {
	now := time.Now()
	_, err := s.db.NewUpdate().
		Model((*ent.MemberDevice)(nil)).
		Set("member_id = ?", item.MemberID).
		Set("platform = ?", item.Platform).
		Set("device_token = ?", item.DeviceToken).
		Set("device_name = ?", item.DeviceName).
		Set("app_version = ?", item.AppVersion).
		Set("is_active = ?", item.IsActive).
		Set("last_seen_at = ?", item.LastSeenAt).
		Set("updated_at = ?", now).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return s.GetMemberDeviceByID(ctx, id)
}

func (s *Service) DeleteMemberDeviceByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.MemberDevice)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
