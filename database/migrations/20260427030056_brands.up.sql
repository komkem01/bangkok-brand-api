SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE brands (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name_th varchar,
	name_en varchar,
	logo_id uuid REFERENCES storages(id),
	description text,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE brands IS 'ตารางเก็บข้อมูลแบรนด์หรือกลุ่มวิสาหกิจชุมชน';
COMMENT ON COLUMN brands.name_th IS 'ชื่อแบรนด์หรือตราสินค้าภาษาไทย';
COMMENT ON COLUMN brands.name_en IS 'ชื่อแบรนด์ภาษาอังกฤษ';
COMMENT ON COLUMN brands.logo_id IS 'โลโก้ประจำแบรนด์';
COMMENT ON COLUMN brands.description IS 'ข้อมูลรายละเอียดแบรนด์';
COMMENT ON COLUMN brands.is_active IS 'สถานะแบรนด์';
