SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_settlement_items_batch_id;
DROP TABLE IF EXISTS settlement_items;
