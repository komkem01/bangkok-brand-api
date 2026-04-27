SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE return_items (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	return_request_id uuid REFERENCES return_requests(id),
	order_item_id uuid REFERENCES order_items(id),
	quantity integer NOT NULL DEFAULT 1,
	reason varchar,
	condition_note text,
	evidence_storage_id uuid REFERENCES storages(id),
	refund_amount decimal(12, 2) NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (return_request_id, order_item_id)
);

--bun:split

COMMENT ON TABLE return_items IS 'ตารางรายการสินค้าที่อยู่ในคำขอคืนแต่ละครั้ง';
COMMENT ON COLUMN return_items.return_request_id IS 'คำขอคืนสินค้าที่รายการนี้สังกัด';
COMMENT ON COLUMN return_items.order_item_id IS 'รายการสินค้าที่ถูกขอคืน';
COMMENT ON COLUMN return_items.quantity IS 'จำนวนชิ้นที่ขอคืน';
COMMENT ON COLUMN return_items.reason IS 'เหตุผลการคืนรายชิ้น';
COMMENT ON COLUMN return_items.condition_note IS 'สภาพสินค้าหรือคำอธิบายเพิ่มเติม';
COMMENT ON COLUMN return_items.evidence_storage_id IS 'ไฟล์รูป/หลักฐานประกอบการคืนสินค้า';
COMMENT ON COLUMN return_items.refund_amount IS 'ยอดเงินคืนของรายการนี้';
