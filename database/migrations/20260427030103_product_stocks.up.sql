SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_stocks (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	product_id uuid REFERENCES products(id),
	quantity integer NOT NULL DEFAULT 0,
	low_stock_threshold integer NOT NULL DEFAULT 5,
	last_restocked_at timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_stocks IS 'ตารางจัดการปริมาณสินค้าคงคลัง';
COMMENT ON COLUMN product_stocks.product_id IS 'สินค้าที่ต้องการจัดการสต็อก';
COMMENT ON COLUMN product_stocks.quantity IS 'จำนวนสินค้าที่มีอยู่ในคลังปัจจุบัน';
COMMENT ON COLUMN product_stocks.low_stock_threshold IS 'จำนวนขั้นต่ำที่ต้องแจ้งเตือนเมื่อสินค้าใกล้หมด';
COMMENT ON COLUMN product_stocks.last_restocked_at IS 'วันเวลาที่มีการเติมสต็อกครั้งล่าสุด';
