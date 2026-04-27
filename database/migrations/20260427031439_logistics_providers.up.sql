SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE logistics_providers (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name_th varchar,
	name_en varchar,
	code varchar UNIQUE,
	tracking_url_template text,
	api_endpoint text,
	supports_cod boolean NOT NULL DEFAULT false,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE logistics_providers IS 'ตารางเก็บข้อมูลผู้ให้บริการขนส่ง';
COMMENT ON COLUMN logistics_providers.code IS 'รหัสผู้ให้บริการขนส่ง';
COMMENT ON COLUMN logistics_providers.tracking_url_template IS 'รูปแบบ URL สำหรับติดตามพัสดุ';
COMMENT ON COLUMN logistics_providers.api_endpoint IS 'ปลายทาง API ของผู้ให้บริการขนส่ง';
COMMENT ON COLUMN logistics_providers.supports_cod IS 'รองรับการเก็บเงินปลายทางหรือไม่';
