SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_flash_sale_events_active_window;
DROP INDEX IF EXISTS idx_flash_sale_events_status;
DROP INDEX IF EXISTS idx_flash_sale_events_shop_id;
DROP TABLE IF EXISTS flash_sale_events;

--bun:split

DROP TYPE IF EXISTS flash_sale_status;
