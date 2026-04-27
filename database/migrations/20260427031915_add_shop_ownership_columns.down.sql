SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_payments_shop_id;
COMMENT ON COLUMN payments.shop_id IS NULL;
ALTER TABLE payments DROP COLUMN IF EXISTS shop_id;

--bun:split

DROP INDEX IF EXISTS idx_rewards_shop_id;
COMMENT ON COLUMN rewards.shop_id IS NULL;
ALTER TABLE rewards DROP COLUMN IF EXISTS shop_id;

--bun:split

DROP INDEX IF EXISTS idx_coupons_shop_id;
COMMENT ON COLUMN coupons.shop_id IS NULL;
ALTER TABLE coupons DROP COLUMN IF EXISTS shop_id;

--bun:split

DROP INDEX IF EXISTS idx_orders_shop_id;
COMMENT ON COLUMN orders.shop_id IS NULL;
ALTER TABLE orders DROP COLUMN IF EXISTS shop_id;

--bun:split

DROP INDEX IF EXISTS idx_products_shop_id;
COMMENT ON COLUMN products.shop_id IS NULL;
ALTER TABLE products DROP COLUMN IF EXISTS shop_id;
