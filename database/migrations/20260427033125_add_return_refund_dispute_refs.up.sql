SET statement_timeout = 0;

--bun:split

ALTER TABLE orders
	ADD COLUMN IF NOT EXISTS latest_return_request_id uuid REFERENCES return_requests(id),
	ADD COLUMN IF NOT EXISTS has_open_dispute boolean NOT NULL DEFAULT false,
	ADD COLUMN IF NOT EXISTS refunded_amount decimal(12, 2) NOT NULL DEFAULT 0;

COMMENT ON COLUMN orders.latest_return_request_id IS 'คำขอคืนสินค้าล่าสุดของออเดอร์';
COMMENT ON COLUMN orders.has_open_dispute IS 'ออเดอร์นี้มีเคสข้อพิพาทที่ยังไม่ปิดหรือไม่';
COMMENT ON COLUMN orders.refunded_amount IS 'ยอดเงินคืนสะสมของออเดอร์';

--bun:split

ALTER TABLE payments
	ADD COLUMN IF NOT EXISTS refunded_amount decimal(12, 2) NOT NULL DEFAULT 0,
	ADD COLUMN IF NOT EXISTS last_refund_at timestamp;

COMMENT ON COLUMN payments.refunded_amount IS 'ยอดเงินคืนสะสมของรายการชำระเงิน';
COMMENT ON COLUMN payments.last_refund_at IS 'วันเวลาที่มีการคืนเงินล่าสุด';

--bun:split

ALTER TABLE order_items
	ADD COLUMN IF NOT EXISTS returned_quantity integer NOT NULL DEFAULT 0;

COMMENT ON COLUMN order_items.returned_quantity IS 'จำนวนชิ้นที่ถูกคืนแล้วของรายการนี้';

--bun:split

SELECT 2
