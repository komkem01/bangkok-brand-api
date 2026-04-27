SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE return_requests (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	request_no varchar UNIQUE,
	order_id uuid REFERENCES orders(id),
	member_id uuid REFERENCES members(id),
	shop_id uuid REFERENCES shops(id),
	status return_request_status NOT NULL DEFAULT 'requested',
	reason varchar,
	detail text,
	requested_at timestamp,
	approved_at timestamp,
	rejected_at timestamp,
	received_at timestamp,
	processed_by_id uuid REFERENCES members(id),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_return_requests_order_id ON return_requests(order_id);
CREATE INDEX IF NOT EXISTS idx_return_requests_member_id ON return_requests(member_id);
CREATE INDEX IF NOT EXISTS idx_return_requests_shop_id ON return_requests(shop_id);

--bun:split

COMMENT ON TABLE return_requests IS 'ตารางคำขอคืนสินค้า/คืนเงินจากลูกค้า';
COMMENT ON COLUMN return_requests.request_no IS 'เลขที่คำขอคืนสินค้า';
COMMENT ON COLUMN return_requests.order_id IS 'คำสั่งซื้อที่ต้องการคืนสินค้า';
COMMENT ON COLUMN return_requests.member_id IS 'ลูกค้าที่ส่งคำขอ';
COMMENT ON COLUMN return_requests.shop_id IS 'ร้านค้าที่เกี่ยวข้องกับคำขอ';
COMMENT ON COLUMN return_requests.status IS 'สถานะปัจจุบันของคำขอ';
COMMENT ON COLUMN return_requests.reason IS 'เหตุผลหลักในการขอคืนสินค้า';
COMMENT ON COLUMN return_requests.detail IS 'รายละเอียดเพิ่มเติมของปัญหา';
COMMENT ON COLUMN return_requests.processed_by_id IS 'ผู้ดูแลที่ดำเนินการเคส';
