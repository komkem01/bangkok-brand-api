SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_variants (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	product_id uuid REFERENCES products(id),
	shop_id uuid REFERENCES shops(id),
	sku varchar UNIQUE,
	name_th varchar,
	name_en varchar,
	barcode varchar,
	price decimal(12, 2),
	discount_price decimal(12, 2),
	additional_price decimal(12, 2) NOT NULL DEFAULT 0,
	is_default boolean NOT NULL DEFAULT false,
	is_active boolean NOT NULL DEFAULT true,
	weight decimal(8, 2),
	width decimal(8, 2),
	length decimal(8, 2),
	height decimal(8, 2),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_variants IS 'ตารางเก็บตัวเลือกย่อยของสินค้า (เช่น สี/ไซซ์)';
COMMENT ON COLUMN product_variants.product_id IS 'สินค้าแม่ของตัวเลือกย่อย';
COMMENT ON COLUMN product_variants.shop_id IS 'ร้านเจ้าของตัวเลือกย่อย';
COMMENT ON COLUMN product_variants.sku IS 'รหัส SKU ระดับตัวเลือกย่อย';
COMMENT ON COLUMN product_variants.barcode IS 'บาร์โค้ดของตัวเลือกย่อย';
COMMENT ON COLUMN product_variants.price IS 'ราคาขายของตัวเลือกย่อย';
COMMENT ON COLUMN product_variants.discount_price IS 'ราคาหลังลดของตัวเลือกย่อย';
COMMENT ON COLUMN product_variants.additional_price IS 'ราคาเพิ่มจากสินค้าหลัก';
COMMENT ON COLUMN product_variants.is_default IS 'กำหนดให้เป็นตัวเลือกเริ่มต้น';
