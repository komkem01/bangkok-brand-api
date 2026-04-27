SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE dispute_cases (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	case_no varchar UNIQUE,
	order_id uuid REFERENCES orders(id),
	return_request_id uuid REFERENCES return_requests(id),
	member_id uuid REFERENCES members(id),
	shop_id uuid REFERENCES shops(id),
	subject varchar,
	detail text,
	status dispute_case_status NOT NULL DEFAULT 'open',
	assigned_admin_id uuid REFERENCES members(id),
	opened_at timestamp,
	closed_at timestamp,
	resolution_note text,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_dispute_cases_order_id ON dispute_cases(order_id);
CREATE INDEX IF NOT EXISTS idx_dispute_cases_shop_id ON dispute_cases(shop_id);

--bun:split

COMMENT ON TABLE dispute_cases IS 'ตารางเคสข้อพิพาทระหว่างลูกค้าและร้านค้า';
COMMENT ON COLUMN dispute_cases.case_no IS 'เลขที่เคสข้อพิพาท';
COMMENT ON COLUMN dispute_cases.order_id IS 'ออเดอร์ที่เป็นประเด็นข้อพิพาท';
COMMENT ON COLUMN dispute_cases.return_request_id IS 'คำขอคืนสินค้าที่เกี่ยวข้อง (ถ้ามี)';
COMMENT ON COLUMN dispute_cases.member_id IS 'ลูกค้าผู้เปิดเคส';
COMMENT ON COLUMN dispute_cases.shop_id IS 'ร้านที่ถูกร้องเรียน';
COMMENT ON COLUMN dispute_cases.status IS 'สถานะปัจจุบันของเคส';
COMMENT ON COLUMN dispute_cases.assigned_admin_id IS 'แอดมินที่รับผิดชอบเคส';
COMMENT ON COLUMN dispute_cases.resolution_note IS 'สรุปผลการตัดสินหรือข้อตกลง';
