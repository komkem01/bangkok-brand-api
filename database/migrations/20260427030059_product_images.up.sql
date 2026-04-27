SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_images (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	product_id uuid REFERENCES products(id),
	storage_id uuid REFERENCES storages(id),
	is_main boolean NOT NULL DEFAULT false,
	sort_order integer NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_images IS 'ตารางเก็บรูปภาพประกอบสินค้า (หนึ่งสินค้ามีได้หลายรูป)';
COMMENT ON COLUMN product_images.product_id IS 'สินค้าที่เป็นเจ้าของรูปภาพนี้';
COMMENT ON COLUMN product_images.storage_id IS 'ไฟล์รูปภาพที่เก็บอยู่ในตาราง storage';
COMMENT ON COLUMN product_images.is_main IS 'กำหนดให้เป็นรูปภาพหน้าปกหลัก';
COMMENT ON COLUMN product_images.sort_order IS 'ลำดับการแสดงผลรูปในแกลเลอรี';
