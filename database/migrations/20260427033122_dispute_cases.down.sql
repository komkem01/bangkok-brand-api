SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_dispute_cases_shop_id;
DROP INDEX IF EXISTS idx_dispute_cases_order_id;
DROP TABLE IF EXISTS dispute_cases;
