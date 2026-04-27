SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE banks (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name_th varchar,
	name_en varchar,
	code varchar,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE banks IS 'ตารางเก็บข้อมูลธนาคาร';
COMMENT ON COLUMN banks.name_th IS 'ชื่อธนาคารภาษาไทย';
COMMENT ON COLUMN banks.name_en IS 'ชื่อธนาคารภาษาอังกฤษ';
COMMENT ON COLUMN banks.code IS 'รหัสธนาคาร';
COMMENT ON COLUMN banks.is_active IS 'สถานะการใช้งานของธนาคาร';
