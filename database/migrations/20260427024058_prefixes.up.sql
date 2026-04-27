SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE prefixes (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	gender_id uuid REFERENCES genders(id),
	name_th varchar,
	name_en varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE prefixes IS 'ตารางเก็บข้อมูลคำนำหน้าชื่อ';
COMMENT ON COLUMN prefixes.gender_id IS 'อ้างอิงเพศของคำนำหน้าชื่อ';
COMMENT ON COLUMN prefixes.name_th IS 'ชื่อคำนำหน้าภาษาไทย';
COMMENT ON COLUMN prefixes.name_en IS 'ชื่อคำนำหน้าภาษาอังกฤษ';
COMMENT ON COLUMN prefixes.is_active IS 'สถานะการใช้งานของข้อมูลคำนำหน้า';
