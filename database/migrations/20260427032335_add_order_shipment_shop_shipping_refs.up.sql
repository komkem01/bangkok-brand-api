SET statement_timeout = 0;

--bun:split

ALTER TABLE order_shipments
	ADD COLUMN IF NOT EXISTS shop_shipping_method_id uuid REFERENCES shop_shipping_methods(id);

COMMENT ON COLUMN order_shipments.shop_shipping_method_id IS 'วิธีจัดส่งรายร้านที่ถูกใช้จริงในออเดอร์';
CREATE INDEX IF NOT EXISTS idx_order_shipments_shop_shipping_method_id ON order_shipments(shop_shipping_method_id);

--bun:split

ALTER TABLE order_shipments
	ADD COLUMN IF NOT EXISTS shipping_zone_id uuid REFERENCES shipping_zones(id);

COMMENT ON COLUMN order_shipments.shipping_zone_id IS 'โซนจัดส่งที่ถูกใช้ตอนคำนวณค่าส่ง';
CREATE INDEX IF NOT EXISTS idx_order_shipments_shipping_zone_id ON order_shipments(shipping_zone_id);
