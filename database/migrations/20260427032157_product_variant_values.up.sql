SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_variant_values (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	variant_id uuid REFERENCES product_variants(id),
	attribute_id uuid REFERENCES product_attributes(id),
	value_th varchar,
	value_en varchar,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (variant_id, attribute_id)
);

--bun:split

COMMENT ON TABLE product_variant_values IS 'ตารางเก็บค่าคุณสมบัติที่ผูกกับตัวเลือกย่อยของสินค้า';
COMMENT ON COLUMN product_variant_values.variant_id IS 'อ้างอิงตัวเลือกย่อยของสินค้า';
COMMENT ON COLUMN product_variant_values.attribute_id IS 'ประเภทคุณสมบัติ เช่น สี/ไซซ์';
COMMENT ON COLUMN product_variant_values.value_th IS 'ค่าคุณสมบัติภาษาไทย';
COMMENT ON COLUMN product_variant_values.value_en IS 'ค่าคุณสมบัติภาษาอังกฤษ';
