SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_refund_transactions_payment_id;
DROP INDEX IF EXISTS idx_refund_transactions_order_id;
DROP TABLE IF EXISTS refund_transactions;
