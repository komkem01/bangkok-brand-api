SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE zipcodes (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	sub_district_id uuid REFERENCES sub_districts(id),
	name varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE zipcodes IS 'ตารางเก็บข้อมูลรหัสไปรษณีย์';
COMMENT ON COLUMN zipcodes.sub_district_id IS 'อ้างอิงตำบลหรือแขวงของรหัสไปรษณีย์';
COMMENT ON COLUMN zipcodes.name IS 'รหัสไปรษณีย์';
COMMENT ON COLUMN zipcodes.is_active IS 'สถานะการใช้งานของข้อมูลรหัสไปรษณีย์';
