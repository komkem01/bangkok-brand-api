SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE carts (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid UNIQUE REFERENCES members(id),
	total_items integer NOT NULL DEFAULT 0,
	total_price decimal(12, 2) NOT NULL DEFAULT 0.00,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE carts IS 'ตารางหลักสำหรับจัดการตะกร้าสินค้าของผู้ใช้งาน';
COMMENT ON COLUMN carts.member_id IS 'สมาชิกที่เป็นเจ้าของตะกร้านี้ (1 สมาชิกมี 1 ตะกร้า)';
COMMENT ON COLUMN carts.total_items IS 'จำนวนรายการสินค้าทั้งหมดในตะกร้า';
COMMENT ON COLUMN carts.total_price IS 'ราคารวมทั้งหมดในตะกร้า';
