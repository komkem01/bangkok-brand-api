SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_return_requests_shop_id;
DROP INDEX IF EXISTS idx_return_requests_member_id;
DROP INDEX IF EXISTS idx_return_requests_order_id;
DROP TABLE IF EXISTS return_requests;
