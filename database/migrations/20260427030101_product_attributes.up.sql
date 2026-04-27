SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_attributes (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name_th varchar,
	name_en varchar,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_attributes IS 'ตารางมาสเตอร์กำหนดประเภทคุณสมบัติพิเศษของสินค้า';
COMMENT ON COLUMN product_attributes.name_th IS 'ชื่อหัวข้อคุณสมบัติ เช่น สี, ขนาด, รสชาติ';
