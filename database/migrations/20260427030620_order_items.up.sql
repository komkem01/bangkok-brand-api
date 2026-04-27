SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE order_items (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	order_id uuid REFERENCES orders(id),
	product_id uuid REFERENCES products(id),
	product_name_snapshot varchar,
	sku_snapshot varchar,
	quantity integer,
	unit_price decimal(12, 2),
	selected_attributes json,
	subtotal_price decimal(12, 2),
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE order_items IS 'ตารางเก็บรายการสินค้าภายใต้คำสั่งซื้อ (Snapshot ข้อมูลเพื่อป้องกันการเปลี่ยนแปลงภายหลัง)';
COMMENT ON COLUMN order_items.order_id IS 'คำสั่งซื้อที่รายการนี้สังกัด';
COMMENT ON COLUMN order_items.product_id IS 'สินค้าที่ซื้อ';
COMMENT ON COLUMN order_items.product_name_snapshot IS 'ชื่อสินค้า ณ เวลาที่ซื้อ';
COMMENT ON COLUMN order_items.sku_snapshot IS 'SKU ณ เวลาที่ซื้อ';
COMMENT ON COLUMN order_items.quantity IS 'จำนวนที่ซื้อ';
COMMENT ON COLUMN order_items.unit_price IS 'ราคาต่อหน่วย ณ เวลาที่ซื้อ';
COMMENT ON COLUMN order_items.selected_attributes IS 'คุณสมบัติที่เลือก ณ เวลาที่ซื้อ';
COMMENT ON COLUMN order_items.subtotal_price IS 'ราคารวมรายการนี้';
