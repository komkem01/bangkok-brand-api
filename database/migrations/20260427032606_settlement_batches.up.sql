SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE settlement_batches (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid REFERENCES shops(id),
	status settlement_status NOT NULL DEFAULT 'pending',
	period_start timestamp,
	period_end timestamp,
	gross_amount decimal(12, 2) NOT NULL DEFAULT 0,
	fee_amount decimal(12, 2) NOT NULL DEFAULT 0,
	adjustment_amount decimal(12, 2) NOT NULL DEFAULT 0,
	net_amount decimal(12, 2) NOT NULL DEFAULT 0,
	transfer_fee decimal(12, 2) NOT NULL DEFAULT 0,
	final_amount decimal(12, 2) NOT NULL DEFAULT 0,
	payout_account_id uuid REFERENCES member_bank_accounts(id),
	requested_by_id uuid REFERENCES members(id),
	approved_by_id uuid REFERENCES members(id),
	paid_by_id uuid REFERENCES members(id),
	payout_reference varchar,
	paid_at timestamp,
	note text,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_settlement_batches_shop_id ON settlement_batches(shop_id);

--bun:split

COMMENT ON TABLE settlement_batches IS 'ตารางรอบสรุปรายได้และโอนเงินให้ร้านค้า';
COMMENT ON COLUMN settlement_batches.shop_id IS 'ร้านค้าที่รับเงินรอบนี้';
COMMENT ON COLUMN settlement_batches.status IS 'สถานะของรอบโอนเงิน';
COMMENT ON COLUMN settlement_batches.period_start IS 'วันเริ่มต้นรอบคำนวณ';
COMMENT ON COLUMN settlement_batches.period_end IS 'วันสิ้นสุดรอบคำนวณ';
COMMENT ON COLUMN settlement_batches.gross_amount IS 'ยอดขายรวมก่อนหักค่าธรรมเนียม';
COMMENT ON COLUMN settlement_batches.fee_amount IS 'ค่าธรรมเนียมแพลตฟอร์มรวม';
COMMENT ON COLUMN settlement_batches.adjustment_amount IS 'ยอดปรับเพิ่ม/ลด';
COMMENT ON COLUMN settlement_batches.net_amount IS 'ยอดสุทธิหลังหักค่าธรรมเนียม';
COMMENT ON COLUMN settlement_batches.final_amount IS 'ยอดสุดท้ายที่ต้องโอนให้ร้าน';
COMMENT ON COLUMN settlement_batches.payout_account_id IS 'บัญชีธนาคารปลายทางของร้าน';
