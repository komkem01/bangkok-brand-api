SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE sub_districts (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	district_id uuid REFERENCES districts(id),
	name varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE sub_districts IS 'ตารางเก็บข้อมูลตำบลหรือแขวง';
COMMENT ON COLUMN sub_districts.district_id IS 'อ้างอิงอำเภอหรือเขตของตำบลหรือแขวง';
COMMENT ON COLUMN sub_districts.name IS 'ชื่อตำบลหรือแขวง';
COMMENT ON COLUMN sub_districts.is_active IS 'สถานะการใช้งานของข้อมูลตำบลหรือแขวง';
