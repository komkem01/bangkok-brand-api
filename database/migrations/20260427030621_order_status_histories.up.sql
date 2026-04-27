SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE order_status_histories (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	order_id uuid REFERENCES orders(id),
	status order_status,
	remark text,
	changed_by_id uuid,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE order_status_histories IS 'ตารางบันทึกประวัติการเปลี่ยนแปลงสถานะคำสั่งซื้อ';
COMMENT ON COLUMN order_status_histories.order_id IS 'อ้างอิงคำสั่งซื้อ';
COMMENT ON COLUMN order_status_histories.status IS 'สถานะที่เปลี่ยนไป';
COMMENT ON COLUMN order_status_histories.remark IS 'เหตุผลหรือข้อมูลเพิ่มเติมในการเปลี่ยนสถานะ';
COMMENT ON COLUMN order_status_histories.changed_by_id IS 'ID ของผู้ที่ทำการเปลี่ยนสถานะ (Member หรือ Admin)';
