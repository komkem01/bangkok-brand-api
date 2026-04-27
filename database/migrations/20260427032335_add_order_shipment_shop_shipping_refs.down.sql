SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_order_shipments_shipping_zone_id;
COMMENT ON COLUMN order_shipments.shipping_zone_id IS NULL;
ALTER TABLE order_shipments DROP COLUMN IF EXISTS shipping_zone_id;

--bun:split

DROP INDEX IF EXISTS idx_order_shipments_shop_shipping_method_id;
COMMENT ON COLUMN order_shipments.shop_shipping_method_id IS NULL;
ALTER TABLE order_shipments DROP COLUMN IF EXISTS shop_shipping_method_id;
