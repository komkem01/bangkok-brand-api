SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_flash_sale_items_product_id;
DROP INDEX IF EXISTS idx_flash_sale_items_event_id;
DROP TABLE IF EXISTS flash_sale_items;
