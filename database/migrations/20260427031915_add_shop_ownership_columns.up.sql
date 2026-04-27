SET statement_timeout = 0;

--bun:split

ALTER TABLE products
	ADD COLUMN IF NOT EXISTS shop_id uuid REFERENCES shops(id);

COMMENT ON COLUMN products.shop_id IS 'ร้านค้าเจ้าของสินค้า';
CREATE INDEX IF NOT EXISTS idx_products_shop_id ON products(shop_id);

--bun:split

ALTER TABLE orders
	ADD COLUMN IF NOT EXISTS shop_id uuid REFERENCES shops(id);

COMMENT ON COLUMN orders.shop_id IS 'ร้านค้าที่รับคำสั่งซื้อ';
CREATE INDEX IF NOT EXISTS idx_orders_shop_id ON orders(shop_id);

--bun:split

ALTER TABLE coupons
	ADD COLUMN IF NOT EXISTS shop_id uuid REFERENCES shops(id);

COMMENT ON COLUMN coupons.shop_id IS 'ร้านเจ้าของคูปอง (ถ้าเป็นคูปองเฉพาะร้าน)';
CREATE INDEX IF NOT EXISTS idx_coupons_shop_id ON coupons(shop_id);

--bun:split

ALTER TABLE rewards
	ADD COLUMN IF NOT EXISTS shop_id uuid REFERENCES shops(id);

COMMENT ON COLUMN rewards.shop_id IS 'ร้านเจ้าของของรางวัล';
CREATE INDEX IF NOT EXISTS idx_rewards_shop_id ON rewards(shop_id);

--bun:split

ALTER TABLE payments
	ADD COLUMN IF NOT EXISTS shop_id uuid REFERENCES shops(id);

COMMENT ON COLUMN payments.shop_id IS 'ร้านค้าที่เกี่ยวข้องกับรายการชำระเงิน';
CREATE INDEX IF NOT EXISTS idx_payments_shop_id ON payments(shop_id);
