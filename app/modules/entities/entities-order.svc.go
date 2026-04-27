package entities

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"
	entitiesinf "bangkok-brand/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.OrderEntity = (*Service)(nil)

func (s *Service) ListOrders(ctx context.Context) ([]*ent.Order, error) {
	var items []*ent.Order
	err := s.db.NewSelect().Model(&items).OrderExpr("created_at DESC").Scan(ctx)
	return items, err
}

func (s *Service) GetOrderByID(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	item := &ent.Order{}
	err := s.db.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Service) CreateOrder(ctx context.Context, p *ent.Order) (*ent.Order, error) {
	_, err := s.db.NewInsert().Model(p).Returning("*").Exec(ctx)
	return p, err
}

func (s *Service) UpdateOrderByID(ctx context.Context, id uuid.UUID, p *ent.Order) (*ent.Order, error) {
	_, err := s.db.NewUpdate().
		Model((*ent.Order)(nil)).
		Set("order_no = ?", p.OrderNo).
		Set("member_id = ?", p.MemberID).
		Set("total_product_price = ?", p.TotalProductPrice).
		Set("shipping_fee = ?", p.ShippingFee).
		Set("discount_amount = ?", p.DiscountAmount).
		Set("net_amount = ?", p.NetAmount).
		Set("recipient_name = ?", p.RecipientName).
		Set("recipient_phone = ?", p.RecipientPhone).
		Set("shipping_address_detail = ?", p.ShippingAddressDetail).
		Set("province_id = ?", p.ProvinceID).
		Set("district_id = ?", p.DistrictID).
		Set("sub_district_id = ?", p.SubDistrictID).
		Set("zipcode_id = ?", p.ZipcodeID).
		Set("status = ?", p.Status).
		Set("tracking_number = ?", p.TrackingNumber).
		Set("courier_name = ?", p.CourierName).
		Set("remark = ?", p.Remark).
		Set("ordered_at = ?", p.OrderedAt).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetOrderByID(ctx, id)
}

func (s *Service) DeleteOrderByID(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model((*ent.Order)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
