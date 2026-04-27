SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE payments (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	order_id uuid REFERENCES orders(id),
	payment_no varchar UNIQUE,
	method payment_method,
	amount decimal(12, 2),
	status payment_status NOT NULL DEFAULT 'pending',
	evidence_storage_id uuid REFERENCES storages(id),
	transfer_date_time timestamp,
	from_bank_id uuid REFERENCES banks(id),
	transaction_ref varchar,
	paid_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE payments IS 'ตารางเก็บข้อมูลการชำระเงินและหลักฐานการโอน';
COMMENT ON COLUMN payments.order_id IS 'คำสั่งซื้อที่ชำระเงิน';
COMMENT ON COLUMN payments.payment_no IS 'หมายเลขรายการชำระเงิน';
COMMENT ON COLUMN payments.method IS 'ช่องทางการชำระเงิน';
COMMENT ON COLUMN payments.amount IS 'จำนวนเงินที่ชำระจริง';
COMMENT ON COLUMN payments.status IS 'สถานะการชำระเงิน';
COMMENT ON COLUMN payments.evidence_storage_id IS 'สลิปการโอนเงินที่เก็บในตาราง storage';
COMMENT ON COLUMN payments.transfer_date_time IS 'วันเวลาที่โอนตามสลิป';
COMMENT ON COLUMN payments.from_bank_id IS 'ธนาคารต้นทาง';
COMMENT ON COLUMN payments.transaction_ref IS 'เลขที่อ้างอิงจากระบบ Payment Gateway (ถ้ามี)';
COMMENT ON COLUMN payments.paid_at IS 'วันเวลาที่ยืนยันการชำระเงิน';
