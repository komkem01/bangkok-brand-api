SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE coupon_usages (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	coupon_id uuid REFERENCES coupons(id),
	member_id uuid REFERENCES members(id),
	order_id uuid REFERENCES orders(id),
	discount_applied decimal(12, 2),
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE coupon_usages IS 'ตารางบันทึกประวัติการใช้คูปองของสมาชิก';
COMMENT ON COLUMN coupon_usages.coupon_id IS 'อ้างอิงคูปองที่ถูกใช้งาน';
COMMENT ON COLUMN coupon_usages.member_id IS 'อ้างอิงสมาชิกที่ใช้คูปอง';
COMMENT ON COLUMN coupon_usages.order_id IS 'คำสั่งซื้อที่ใช้คูปองนี้';
COMMENT ON COLUMN coupon_usages.discount_applied IS 'จำนวนส่วนลดที่ได้จริงในออเดอร์นี้';
