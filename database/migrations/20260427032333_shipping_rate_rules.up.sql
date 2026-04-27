SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shipping_rate_rules (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_shipping_method_id uuid REFERENCES shop_shipping_methods(id),
	min_weight decimal(12, 2),
	max_weight decimal(12, 2),
	min_order_amount decimal(12, 2),
	max_order_amount decimal(12, 2),
	rate_amount decimal(12, 2) NOT NULL DEFAULT 0,
	priority integer NOT NULL DEFAULT 1,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE shipping_rate_rules IS 'ตารางกฎคำนวณค่าส่งตามเงื่อนไขน้ำหนักหรือยอดสั่งซื้อ';
COMMENT ON COLUMN shipping_rate_rules.shop_shipping_method_id IS 'อ้างอิงวิธีจัดส่งรายร้าน';
COMMENT ON COLUMN shipping_rate_rules.min_weight IS 'น้ำหนักขั้นต่ำของช่วงกฎ';
COMMENT ON COLUMN shipping_rate_rules.max_weight IS 'น้ำหนักสูงสุดของช่วงกฎ';
COMMENT ON COLUMN shipping_rate_rules.min_order_amount IS 'ยอดสั่งซื้อขั้นต่ำของช่วงกฎ';
COMMENT ON COLUMN shipping_rate_rules.max_order_amount IS 'ยอดสั่งซื้อสูงสุดของช่วงกฎ';
COMMENT ON COLUMN shipping_rate_rules.rate_amount IS 'ค่าจัดส่งที่ต้องคิดตามกฎ';
COMMENT ON COLUMN shipping_rate_rules.priority IS 'ลำดับความสำคัญเมื่อตรงหลายกฎ';
