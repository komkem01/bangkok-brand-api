SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE settlement_items (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	batch_id uuid REFERENCES settlement_batches(id),
	shop_id uuid REFERENCES shops(id),
	order_id uuid REFERENCES orders(id),
	payment_id uuid REFERENCES payments(id),
	status settlement_item_status NOT NULL DEFAULT 'pending',
	gross_amount decimal(12, 2) NOT NULL DEFAULT 0,
	fee_amount decimal(12, 2) NOT NULL DEFAULT 0,
	net_amount decimal(12, 2) NOT NULL DEFAULT 0,
	note text,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_settlement_items_batch_id ON settlement_items(batch_id);

--bun:split

COMMENT ON TABLE settlement_items IS 'ตารางรายการย่อยที่ถูกรวมในแต่ละรอบโอนเงิน';
COMMENT ON COLUMN settlement_items.batch_id IS 'รอบโอนเงินที่รายการนี้สังกัด';
COMMENT ON COLUMN settlement_items.shop_id IS 'ร้านเจ้าของรายการ';
COMMENT ON COLUMN settlement_items.status IS 'สถานะของรายการย่อย';
COMMENT ON COLUMN settlement_items.gross_amount IS 'ยอดก่อนหักค่าธรรมเนียมของรายการย่อย';
COMMENT ON COLUMN settlement_items.fee_amount IS 'ค่าธรรมเนียมของรายการย่อย';
COMMENT ON COLUMN settlement_items.net_amount IS 'ยอดสุทธิของรายการย่อย';
