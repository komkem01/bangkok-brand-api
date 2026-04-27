SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shop_shipping_methods (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid REFERENCES shops(id),
	shipping_method_id uuid REFERENCES shipping_methods(id),
	shipping_zone_id uuid REFERENCES shipping_zones(id),
	method_name varchar,
	fee_adjustment decimal(12, 2) NOT NULL DEFAULT 0,
	free_shipping_threshold decimal(12, 2),
	estimated_days_min integer,
	estimated_days_max integer,
	is_cod_available boolean NOT NULL DEFAULT false,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (shop_id, shipping_method_id, shipping_zone_id)
);

--bun:split

COMMENT ON TABLE shop_shipping_methods IS 'ตารางตั้งค่าวิธีจัดส่งรายร้านและรายโซน';
COMMENT ON COLUMN shop_shipping_methods.shop_id IS 'ร้านเจ้าของการตั้งค่าวิธีจัดส่ง';
COMMENT ON COLUMN shop_shipping_methods.shipping_method_id IS 'อ้างอิงวิธีจัดส่งกลางของระบบ';
COMMENT ON COLUMN shop_shipping_methods.shipping_zone_id IS 'โซนที่วิธีจัดส่งนี้ใช้งาน';
COMMENT ON COLUMN shop_shipping_methods.method_name IS 'ชื่อแสดงผลของวิธีจัดส่งสำหรับร้าน';
COMMENT ON COLUMN shop_shipping_methods.fee_adjustment IS 'ค่าปรับเพิ่ม/ลดจากค่าพื้นฐาน';
COMMENT ON COLUMN shop_shipping_methods.free_shipping_threshold IS 'ยอดขั้นต่ำที่ร้านตั้งค่าส่งฟรี';
COMMENT ON COLUMN shop_shipping_methods.is_cod_available IS 'ร้านรองรับเก็บเงินปลายทางในวิธีนี้';
