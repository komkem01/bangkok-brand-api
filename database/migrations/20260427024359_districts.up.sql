SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE districts (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	province_id uuid REFERENCES provinces(id),
	name varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE districts IS 'ตารางเก็บข้อมูลอำเภอหรือเขต';
COMMENT ON COLUMN districts.province_id IS 'อ้างอิงจังหวัดของอำเภอหรือเขต';
COMMENT ON COLUMN districts.name IS 'ชื่ออำเภอหรือเขต';
COMMENT ON COLUMN districts.is_active IS 'สถานะการใช้งานของข้อมูลอำเภอหรือเขต';
