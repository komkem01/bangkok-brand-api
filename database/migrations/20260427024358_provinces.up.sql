SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE provinces (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE provinces IS 'ตารางเก็บข้อมูลจังหวัด';
COMMENT ON COLUMN provinces.name IS 'ชื่อจังหวัด';
COMMENT ON COLUMN provinces.is_active IS 'สถานะการใช้งานของข้อมูลจังหวัด';
