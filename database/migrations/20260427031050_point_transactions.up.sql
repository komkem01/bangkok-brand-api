SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE point_transactions (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	order_id uuid REFERENCES orders(id),
	type point_transaction_type,
	points integer,
	balance_snapshot integer,
	description text,
	expiry_date timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE point_transactions IS 'ตารางบันทึกประวัติการรับและใช้คะแนนสะสม';
COMMENT ON COLUMN point_transactions.member_id IS 'อ้างอิงสมาชิกที่มีรายการคะแนน';
COMMENT ON COLUMN point_transactions.order_id IS 'อ้างอิงออเดอร์ (ถ้ามี)';
COMMENT ON COLUMN point_transactions.type IS 'ประเภทรายการคะแนน';
COMMENT ON COLUMN point_transactions.points IS 'จำนวนคะแนน (ค่าบวกคือได้ ค่าลบคือใช้)';
COMMENT ON COLUMN point_transactions.balance_snapshot IS 'ยอดคะแนนคงเหลือ ณ เวลานั้น';
COMMENT ON COLUMN point_transactions.description IS 'รายละเอียดรายการ';
COMMENT ON COLUMN point_transactions.expiry_date IS 'วันที่คะแนนส่วนนี้จะหมดอายุ (เฉพาะรายการ earn)';
