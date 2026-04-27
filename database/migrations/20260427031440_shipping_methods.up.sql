SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shipping_methods (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	provider_id uuid REFERENCES logistics_providers(id),
	name_th varchar,
	name_en varchar,
	code varchar,
	method_type shipping_method_type,
	base_fee decimal(12, 2) NOT NULL DEFAULT 0,
	free_shipping_threshold decimal(12, 2),
	estimated_days_min integer,
	estimated_days_max integer,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (provider_id, code)
);

--bun:split

COMMENT ON TABLE shipping_methods IS 'ตารางเก็บช่องทางการจัดส่งที่รองรับในระบบ';
COMMENT ON COLUMN shipping_methods.provider_id IS 'ผู้ให้บริการขนส่งของช่องทางนี้';
COMMENT ON COLUMN shipping_methods.method_type IS 'ประเภทช่องทางการจัดส่ง เช่น ด่วน ธรรมดา';
COMMENT ON COLUMN shipping_methods.base_fee IS 'ค่าจัดส่งพื้นฐาน';
COMMENT ON COLUMN shipping_methods.free_shipping_threshold IS 'ยอดสั่งซื้อขั้นต่ำที่ส่งฟรี';
COMMENT ON COLUMN shipping_methods.estimated_days_min IS 'จำนวนวันจัดส่งขั้นต่ำโดยประมาณ';
COMMENT ON COLUMN shipping_methods.estimated_days_max IS 'จำนวนวันจัดส่งสูงสุดโดยประมาณ';
