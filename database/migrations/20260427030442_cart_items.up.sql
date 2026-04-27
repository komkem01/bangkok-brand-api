SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE cart_items (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	cart_id uuid REFERENCES carts(id),
	product_id uuid REFERENCES products(id),
	quantity integer NOT NULL DEFAULT 1,
	selected_attribute_values json,
	unit_price decimal(12, 2),
	subtotal_price decimal(12, 2),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE cart_items IS 'ตารางเก็บรายการสินค้าแต่ละชิ้นในตะกร้า';
COMMENT ON COLUMN cart_items.cart_id IS 'ตะกร้าที่สินค้าชิ้นนี้สังกัด';
COMMENT ON COLUMN cart_items.product_id IS 'สินค้าที่ถูกเพิ่มเข้ามา';
COMMENT ON COLUMN cart_items.quantity IS 'จำนวนที่ต้องการซื้อ';
COMMENT ON COLUMN cart_items.selected_attribute_values IS 'ข้อมูล JSON เก็บคุณสมบัติที่เลือก เช่น [{"attribute": "color", "value": "Red"}]';
COMMENT ON COLUMN cart_items.unit_price IS 'ราคาต่อหน่วย ณ เวลาที่เพิ่มลงตะกร้า';
COMMENT ON COLUMN cart_items.subtotal_price IS 'ราคารวมของรายการนี้ (quantity * unit_price)';
