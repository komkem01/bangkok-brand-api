SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE order_shipments (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	order_id uuid UNIQUE REFERENCES orders(id),
	provider_id uuid REFERENCES logistics_providers(id),
	shipping_method_id uuid REFERENCES shipping_methods(id),
	tracking_number varchar UNIQUE,
	status shipment_status NOT NULL DEFAULT 'pending_pickup',
	shipping_fee decimal(12, 2) NOT NULL DEFAULT 0,
	receiver_name varchar,
	receiver_phone varchar,
	shipping_address text,
	shipped_at timestamp,
	expected_delivery_at timestamp,
	delivered_at timestamp,
	last_status_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE order_shipments IS 'ตารางขนส่งหลักของคำสั่งซื้อสำหรับติดตามพัสดุ';
COMMENT ON COLUMN order_shipments.order_id IS 'คำสั่งซื้อที่อ้างอิงการจัดส่ง';
COMMENT ON COLUMN order_shipments.provider_id IS 'ผู้ให้บริการขนส่งที่รับงาน';
COMMENT ON COLUMN order_shipments.shipping_method_id IS 'ช่องทางจัดส่งที่เลือก';
COMMENT ON COLUMN order_shipments.tracking_number IS 'หมายเลขติดตามพัสดุ';
COMMENT ON COLUMN order_shipments.status IS 'สถานะปัจจุบันของการจัดส่ง';
COMMENT ON COLUMN order_shipments.shipping_fee IS 'ค่าจัดส่งจริงของคำสั่งซื้อ';
COMMENT ON COLUMN order_shipments.last_status_at IS 'วันเวลาที่อัปเดตสถานะล่าสุด';
