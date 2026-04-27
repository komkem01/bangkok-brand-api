SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE refund_transactions (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	refund_no varchar UNIQUE,
	return_request_id uuid REFERENCES return_requests(id),
	order_id uuid REFERENCES orders(id),
	payment_id uuid REFERENCES payments(id),
	shop_id uuid REFERENCES shops(id),
	status refund_transaction_status NOT NULL DEFAULT 'pending',
	method payment_method,
	amount decimal(12, 2) NOT NULL DEFAULT 0,
	gateway_ref varchar,
	remark text,
	processed_by_id uuid REFERENCES members(id),
	processed_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_refund_transactions_order_id ON refund_transactions(order_id);
CREATE INDEX IF NOT EXISTS idx_refund_transactions_payment_id ON refund_transactions(payment_id);

--bun:split

COMMENT ON TABLE refund_transactions IS 'ตารางบันทึกธุรกรรมคืนเงินให้ลูกค้า';
COMMENT ON COLUMN refund_transactions.refund_no IS 'เลขที่ธุรกรรมคืนเงิน';
COMMENT ON COLUMN refund_transactions.return_request_id IS 'คำขอคืนสินค้าที่นำไปสู่การคืนเงิน';
COMMENT ON COLUMN refund_transactions.order_id IS 'ออเดอร์ที่ถูกคืนเงิน';
COMMENT ON COLUMN refund_transactions.payment_id IS 'รายการชำระเงินต้นทางที่ต้องคืน';
COMMENT ON COLUMN refund_transactions.shop_id IS 'ร้านที่รับผิดชอบการคืนเงิน';
COMMENT ON COLUMN refund_transactions.status IS 'สถานะของธุรกรรมคืนเงิน';
COMMENT ON COLUMN refund_transactions.method IS 'ช่องทางการคืนเงิน';
COMMENT ON COLUMN refund_transactions.amount IS 'ยอดเงินคืนจริง';
COMMENT ON COLUMN refund_transactions.gateway_ref IS 'เลขอ้างอิงจากระบบชำระเงินภายนอก';
