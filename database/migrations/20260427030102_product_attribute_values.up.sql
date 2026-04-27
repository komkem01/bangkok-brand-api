SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_attribute_values (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	product_id uuid REFERENCES products(id),
	attribute_id uuid REFERENCES product_attributes(id),
	value_th varchar,
	value_en varchar,
	additional_price decimal(12, 2) NOT NULL DEFAULT 0
);

--bun:split

COMMENT ON TABLE product_attribute_values IS 'ตารางเก็บค่ารายละเอียดคุณสมบัติเฉพาะรายสินค้า';
COMMENT ON COLUMN product_attribute_values.product_id IS 'สินค้าที่ผูกกับคุณสมบัตินี้';
COMMENT ON COLUMN product_attribute_values.attribute_id IS 'ประเภทคุณสมบัติ';
COMMENT ON COLUMN product_attribute_values.value_th IS 'ค่าของคุณสมบัติ เช่น สีแดง, ไซส์ XL';
COMMENT ON COLUMN product_attribute_values.additional_price IS 'ราคาส่วนต่างที่ต้องบวกเพิ่มจากราคาปกติ (ถ้ามี)';
