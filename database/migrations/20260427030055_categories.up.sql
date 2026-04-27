SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE categories (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	parent_id uuid REFERENCES categories(id),
	name_th varchar,
	name_en varchar,
	description text,
	image_id uuid REFERENCES storages(id),
	slug varchar UNIQUE,
	is_active boolean NOT NULL DEFAULT true,
	sort_order integer NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE categories IS 'ตารางเก็บหมวดหมู่สินค้า รองรับโครงสร้างลำดับชั้น (Hierarchy)';
COMMENT ON COLUMN categories.parent_id IS 'ID ของหมวดหมู่แม่ กรณีเป็นหมวดหมู่ย่อย (Sub-category)';
COMMENT ON COLUMN categories.name_th IS 'ชื่อหมวดหมู่ภาษาไทย';
COMMENT ON COLUMN categories.name_en IS 'ชื่อหมวดหมู่ภาษาอังกฤษ';
COMMENT ON COLUMN categories.description IS 'รายละเอียดคำอธิบายหมวดหมู่';
COMMENT ON COLUMN categories.image_id IS 'รูปภาพประกอบหมวดหมู่';
COMMENT ON COLUMN categories.slug IS 'ชื่อย่อที่ใช้ใน URL เพื่อความ SEO Friendly';
COMMENT ON COLUMN categories.is_active IS 'สถานะการใช้งาน (เปิด/ปิด)';
COMMENT ON COLUMN categories.sort_order IS 'ลำดับการแสดงผลในเมนูหรือหน้าเว็บ';
