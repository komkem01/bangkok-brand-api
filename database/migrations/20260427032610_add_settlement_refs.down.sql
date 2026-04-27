SET statement_timeout = 0;

--bun:split

COMMENT ON COLUMN payments.is_settled IS NULL;
COMMENT ON COLUMN payments.shop_net_amount IS NULL;
COMMENT ON COLUMN payments.platform_fee_amount IS NULL;
ALTER TABLE payments
	DROP COLUMN IF EXISTS is_settled,
	DROP COLUMN IF EXISTS shop_net_amount,
	DROP COLUMN IF EXISTS platform_fee_amount;

--bun:split

DO $$
BEGIN
	IF EXISTS (
		SELECT 1
		FROM information_schema.tables
		WHERE table_schema = current_schema() AND table_name = 'settlement_items'
	) THEN
		DROP INDEX IF EXISTS uq_settlement_items_batch_wallet_transaction;
		COMMENT ON COLUMN settlement_items.wallet_transaction_id IS NULL;
		ALTER TABLE settlement_items DROP COLUMN IF EXISTS wallet_transaction_id;
	END IF;
END
$$;

--bun:split

DROP INDEX IF EXISTS idx_orders_settlement_status;
COMMENT ON COLUMN orders.settlemented_at IS NULL;
COMMENT ON COLUMN orders.settlement_status IS NULL;
COMMENT ON COLUMN orders.is_settlement_eligible IS NULL;
ALTER TABLE orders
	DROP COLUMN IF EXISTS settlemented_at,
	DROP COLUMN IF EXISTS settlement_status,
	DROP COLUMN IF EXISTS is_settlement_eligible;

--bun:split

COMMENT ON COLUMN shops.pending_settlement_amount IS NULL;
COMMENT ON COLUMN shops.available_balance IS NULL;
COMMENT ON COLUMN shops.wallet_balance IS NULL;
ALTER TABLE shops
	DROP COLUMN IF EXISTS pending_settlement_amount,
	DROP COLUMN IF EXISTS available_balance,
	DROP COLUMN IF EXISTS wallet_balance;
