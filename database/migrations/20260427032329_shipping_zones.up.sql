SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shipping_zones (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid REFERENCES shops(id),
	name_th varchar,
	name_en varchar,
	description text,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE shipping_zones IS 'ตารางโซนการจัดส่งรายร้านค้า';
COMMENT ON COLUMN shipping_zones.shop_id IS 'ร้านเจ้าของโซนจัดส่ง';
COMMENT ON COLUMN shipping_zones.name_th IS 'ชื่อโซนจัดส่งภาษาไทย';
COMMENT ON COLUMN shipping_zones.name_en IS 'ชื่อโซนจัดส่งภาษาอังกฤษ';
COMMENT ON COLUMN shipping_zones.description IS 'รายละเอียดพื้นที่ที่ครอบคลุม';
