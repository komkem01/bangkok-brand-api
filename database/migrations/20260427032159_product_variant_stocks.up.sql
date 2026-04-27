SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_variant_stocks (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	variant_id uuid UNIQUE REFERENCES product_variants(id),
	quantity integer NOT NULL DEFAULT 0,
	reserved_quantity integer NOT NULL DEFAULT 0,
	low_stock_threshold integer NOT NULL DEFAULT 5,
	last_restocked_at timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_variant_stocks IS 'ตารางสต็อกระดับตัวเลือกย่อยของสินค้า';
COMMENT ON COLUMN product_variant_stocks.variant_id IS 'อ้างอิงตัวเลือกย่อยของสินค้า';
COMMENT ON COLUMN product_variant_stocks.quantity IS 'จำนวนคงเหลือพร้อมขาย';
COMMENT ON COLUMN product_variant_stocks.reserved_quantity IS 'จำนวนที่ถูกจองในออเดอร์ที่ยังไม่ปิด';
COMMENT ON COLUMN product_variant_stocks.low_stock_threshold IS 'ระดับเตือนสินค้าใกล้หมด';
