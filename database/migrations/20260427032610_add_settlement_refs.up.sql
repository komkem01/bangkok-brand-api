SET statement_timeout = 0;

--bun:split

ALTER TABLE shops
	ADD COLUMN IF NOT EXISTS wallet_balance decimal(12, 2) NOT NULL DEFAULT 0,
	ADD COLUMN IF NOT EXISTS available_balance decimal(12, 2) NOT NULL DEFAULT 0,
	ADD COLUMN IF NOT EXISTS pending_settlement_amount decimal(12, 2) NOT NULL DEFAULT 0;

COMMENT ON COLUMN shops.wallet_balance IS 'ยอดเงินรวมในกระเป๋าร้าน';
COMMENT ON COLUMN shops.available_balance IS 'ยอดเงินที่พร้อมโอนให้ร้าน';
COMMENT ON COLUMN shops.pending_settlement_amount IS 'ยอดเงินที่อยู่ระหว่างรอบโอน';

--bun:split

ALTER TABLE orders
	ADD COLUMN IF NOT EXISTS is_settlement_eligible boolean NOT NULL DEFAULT true,
	ADD COLUMN IF NOT EXISTS settlement_status settlement_status NOT NULL DEFAULT 'pending',
	ADD COLUMN IF NOT EXISTS settlemented_at timestamp;

COMMENT ON COLUMN orders.is_settlement_eligible IS 'ออเดอร์นี้เข้าเงื่อนไขนำไปคำนวณโอนเงินหรือไม่';
COMMENT ON COLUMN orders.settlement_status IS 'สถานะการนำออเดอร์เข้ารอบโอนเงิน';
COMMENT ON COLUMN orders.settlemented_at IS 'วันเวลาที่ออเดอร์นี้ถูกโอนเงินเรียบร้อย';
CREATE INDEX IF NOT EXISTS idx_orders_settlement_status ON orders(settlement_status);

--bun:split

ALTER TABLE payments
	ADD COLUMN IF NOT EXISTS platform_fee_amount decimal(12, 2) NOT NULL DEFAULT 0,
	ADD COLUMN IF NOT EXISTS shop_net_amount decimal(12, 2) NOT NULL DEFAULT 0,
	ADD COLUMN IF NOT EXISTS is_settled boolean NOT NULL DEFAULT false;

COMMENT ON COLUMN payments.platform_fee_amount IS 'ค่าธรรมเนียมแพลตฟอร์มที่หักจากการชำระเงิน';
COMMENT ON COLUMN payments.shop_net_amount IS 'ยอดสุทธิที่เป็นรายได้ของร้าน';
COMMENT ON COLUMN payments.is_settled IS 'การชำระเงินรายการนี้ถูกปิดบัญชีโอนเงินแล้วหรือไม่';

--bun:split

DO $$
BEGIN
	IF EXISTS (
		SELECT 1
		FROM information_schema.tables
		WHERE table_schema = current_schema() AND table_name = 'settlement_items'
	) THEN
		ALTER TABLE settlement_items
			ADD COLUMN IF NOT EXISTS wallet_transaction_id uuid REFERENCES shop_wallet_transactions(id);

		COMMENT ON COLUMN settlement_items.wallet_transaction_id IS 'รายการ ledger ที่ถูกนำมาคำนวณ';

		CREATE UNIQUE INDEX IF NOT EXISTS uq_settlement_items_batch_wallet_transaction
			ON settlement_items(batch_id, wallet_transaction_id)
			WHERE wallet_transaction_id IS NOT NULL;
	END IF;
END
$$;
