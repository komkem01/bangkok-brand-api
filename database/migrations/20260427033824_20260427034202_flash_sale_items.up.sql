SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE flash_sale_items (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	flash_sale_event_id uuid REFERENCES flash_sale_events(id),
	product_id uuid REFERENCES products(id),
	product_variant_id uuid,
	original_price decimal(12, 2) NOT NULL DEFAULT 0,
	sale_price decimal(12, 2) NOT NULL DEFAULT 0,
	discount_percent decimal(5, 2),
	stock_quota integer,
	sold_count integer NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (flash_sale_event_id, product_id, product_variant_id)
);

CREATE INDEX IF NOT EXISTS idx_flash_sale_items_event_id ON flash_sale_items(flash_sale_event_id);
CREATE INDEX IF NOT EXISTS idx_flash_sale_items_product_id ON flash_sale_items(product_id);

--bun:split

COMMENT ON TABLE flash_sale_items IS 'ตารางสินค้าแต่ละรายการที่เข้าร่วมใน flash sale';
COMMENT ON COLUMN flash_sale_items.flash_sale_event_id IS 'เอเวนต์ที่รายการนี้สังกัด';
COMMENT ON COLUMN flash_sale_items.product_id IS 'สินค้าที่ร่วมลดราคา';
COMMENT ON COLUMN flash_sale_items.product_variant_id IS 'ตัวเลือก variant ที่ลดราคา (null = ทุก variant)';
COMMENT ON COLUMN flash_sale_items.original_price IS 'ราคาอ้างอิงก่อนลด';
COMMENT ON COLUMN flash_sale_items.sale_price IS 'ราคาช่วง flash sale';
COMMENT ON COLUMN flash_sale_items.discount_percent IS 'เปอร์เซ็นต์ส่วนลดสำหรับแสดงผล';
COMMENT ON COLUMN flash_sale_items.stock_quota IS 'จำนวนสต็อคที่เปิดขายในเอเวนต์นี้ (null = ไม่จำกัด)';
COMMENT ON COLUMN flash_sale_items.sold_count IS 'จำนวนที่ขายไปแล้วในเอเวนต์นี้';
