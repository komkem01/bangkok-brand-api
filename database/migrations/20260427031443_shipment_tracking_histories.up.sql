SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shipment_tracking_histories (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shipment_id uuid REFERENCES order_shipments(id),
	status shipment_status,
	location varchar,
	description text,
	event_at timestamp,
	raw_payload json,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE shipment_tracking_histories IS 'ตารางบันทึกเหตุการณ์ติดตามสถานะพัสดุ';
COMMENT ON COLUMN shipment_tracking_histories.shipment_id IS 'อ้างอิงรายการจัดส่ง';
COMMENT ON COLUMN shipment_tracking_histories.status IS 'สถานะการจัดส่งในเหตุการณ์นั้น';
COMMENT ON COLUMN shipment_tracking_histories.location IS 'ตำแหน่งหรือศูนย์กระจายสินค้า';
COMMENT ON COLUMN shipment_tracking_histories.description IS 'รายละเอียดเหตุการณ์การติดตาม';
COMMENT ON COLUMN shipment_tracking_histories.event_at IS 'วันเวลาที่เกิดเหตุการณ์จากผู้ให้บริการ';
COMMENT ON COLUMN shipment_tracking_histories.raw_payload IS 'ข้อมูลดิบจากผู้ให้บริการขนส่ง (ถ้ามี)';
