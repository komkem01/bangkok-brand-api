SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_product_images_variant_id;
COMMENT ON COLUMN product_images.variant_id IS NULL;
ALTER TABLE product_images DROP COLUMN IF EXISTS variant_id;

--bun:split

DROP INDEX IF EXISTS idx_order_items_variant_id;
COMMENT ON COLUMN order_items.variant_id IS NULL;
ALTER TABLE order_items DROP COLUMN IF EXISTS variant_id;

--bun:split

DROP INDEX IF EXISTS idx_cart_items_variant_id;
COMMENT ON COLUMN cart_items.variant_id IS NULL;
ALTER TABLE cart_items DROP COLUMN IF EXISTS variant_id;
