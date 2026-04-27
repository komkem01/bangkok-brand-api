SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_settlement_batches_shop_id;
DROP TABLE IF EXISTS settlement_batches;
