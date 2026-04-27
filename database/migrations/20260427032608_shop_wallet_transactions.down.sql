SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_shop_wallet_transactions_created_at;
DROP INDEX IF EXISTS idx_shop_wallet_transactions_shop_id;
DROP TABLE IF EXISTS shop_wallet_transactions;
