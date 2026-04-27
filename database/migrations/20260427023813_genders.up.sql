SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE genders (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name_th varchar,
	name_en varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE genders IS 'ตารางเก็บข้อมูลเพศ';
COMMENT ON COLUMN genders.name_th IS 'ชื่อเพศภาษาไทย';
COMMENT ON COLUMN genders.name_en IS 'ชื่อเพศภาษาอังกฤษ';
COMMENT ON COLUMN genders.is_active IS 'สถานะการใช้งานของข้อมูลเพศ';


