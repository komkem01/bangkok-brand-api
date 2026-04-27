SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shop_wallet_transactions (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid REFERENCES shops(id),
	order_id uuid REFERENCES orders(id),
	payment_id uuid REFERENCES payments(id),
	settlement_batch_id uuid REFERENCES settlement_batches(id),
	tx_type wallet_transaction_type,
	amount decimal(12, 2) NOT NULL DEFAULT 0,
	balance_snapshot decimal(12, 2),
	available_balance_snapshot decimal(12, 2),
	description text,
	metadata json,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_shop_wallet_transactions_shop_id ON shop_wallet_transactions(shop_id);
CREATE INDEX IF NOT EXISTS idx_shop_wallet_transactions_created_at ON shop_wallet_transactions(created_at);

--bun:split

COMMENT ON TABLE shop_wallet_transactions IS 'ตารางสมุดบัญชีรายรับรายจ่ายของร้านค้า (Wallet Ledger)';
COMMENT ON COLUMN shop_wallet_transactions.shop_id IS 'ร้านเจ้าของรายการทางการเงิน';
COMMENT ON COLUMN shop_wallet_transactions.order_id IS 'ออเดอร์ที่เกี่ยวข้อง (ถ้ามี)';
COMMENT ON COLUMN shop_wallet_transactions.payment_id IS 'การชำระเงินที่เกี่ยวข้อง (ถ้ามี)';
COMMENT ON COLUMN shop_wallet_transactions.settlement_batch_id IS 'รอบโอนเงินที่เกี่ยวข้อง (ถ้ามี)';
COMMENT ON COLUMN shop_wallet_transactions.tx_type IS 'ประเภทรายการกระเป๋าเงิน';
COMMENT ON COLUMN shop_wallet_transactions.amount IS 'จำนวนเงิน (+/-) ของรายการ';
COMMENT ON COLUMN shop_wallet_transactions.balance_snapshot IS 'ยอดเงินคงเหลือรวม ณ เวลานั้น';
COMMENT ON COLUMN shop_wallet_transactions.available_balance_snapshot IS 'ยอดเงินพร้อมโอน ณ เวลานั้น';
