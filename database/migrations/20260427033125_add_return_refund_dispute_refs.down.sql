SET statement_timeout = 0;

--bun:split

COMMENT ON COLUMN order_items.returned_quantity IS NULL;
ALTER TABLE order_items DROP COLUMN IF EXISTS returned_quantity;

--bun:split

COMMENT ON COLUMN payments.last_refund_at IS NULL;
COMMENT ON COLUMN payments.refunded_amount IS NULL;
ALTER TABLE payments
	DROP COLUMN IF EXISTS last_refund_at,
	DROP COLUMN IF EXISTS refunded_amount;

--bun:split

COMMENT ON COLUMN orders.refunded_amount IS NULL;
COMMENT ON COLUMN orders.has_open_dispute IS NULL;
COMMENT ON COLUMN orders.latest_return_request_id IS NULL;
ALTER TABLE orders
	DROP COLUMN IF EXISTS refunded_amount,
	DROP COLUMN IF EXISTS has_open_dispute,
	DROP COLUMN IF EXISTS latest_return_request_id;
