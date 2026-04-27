SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE orders (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	order_no varchar UNIQUE,
	member_id uuid REFERENCES members(id),
	total_product_price decimal(12, 2),
	shipping_fee decimal(12, 2) NOT NULL DEFAULT 0.00,
	discount_amount decimal(12, 2) NOT NULL DEFAULT 0.00,
	net_amount decimal(12, 2),
	recipient_name varchar,
	recipient_phone varchar,
	shipping_address_detail text,
	province_id uuid REFERENCES provinces(id),
	district_id uuid REFERENCES districts(id),
	sub_district_id uuid REFERENCES sub_districts(id),
	zipcode_id uuid REFERENCES zipcodes(id),
	status order_status NOT NULL DEFAULT 'pending_payment',
	tracking_number varchar,
	courier_name varchar,
	remark text,
	ordered_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE orders IS 'ตารางเก็บข้อมูลหลักของคำสั่งซื้อ';
COMMENT ON COLUMN orders.order_no IS 'หมายเลขคำสั่งซื้อ (เช่น BK-20240001)';
COMMENT ON COLUMN orders.member_id IS 'ผู้สั่งซื้อ';
COMMENT ON COLUMN orders.total_product_price IS 'ยอดรวมราคาสินค้าทั้งหมด';
COMMENT ON COLUMN orders.shipping_fee IS 'ค่าจัดส่ง';
COMMENT ON COLUMN orders.discount_amount IS 'ยอดส่วนลดรวม (ถ้ามี)';
COMMENT ON COLUMN orders.net_amount IS 'ยอดรวมสุทธิที่ต้องชำระ (Total - Discount + Shipping)';
COMMENT ON COLUMN orders.recipient_name IS 'ชื่อผู้รับ';
COMMENT ON COLUMN orders.recipient_phone IS 'เบอร์ติดต่อผู้รับ';
COMMENT ON COLUMN orders.shipping_address_detail IS 'ที่อยู่จัดส่งโดยละเอียด';
COMMENT ON COLUMN orders.province_id IS 'จังหวัดปลายทางจัดส่ง';
COMMENT ON COLUMN orders.district_id IS 'อำเภอหรือเขตปลายทางจัดส่ง';
COMMENT ON COLUMN orders.sub_district_id IS 'ตำบลหรือแขวงปลายทางจัดส่ง';
COMMENT ON COLUMN orders.zipcode_id IS 'รหัสไปรษณีย์ปลายทางจัดส่ง';
COMMENT ON COLUMN orders.status IS 'สถานะปัจจุบันของคำสั่งซื้อ';
COMMENT ON COLUMN orders.tracking_number IS 'เลขพัสดุสำหรับติดตามสินค้า';
COMMENT ON COLUMN orders.courier_name IS 'บริษัทขนส่ง เช่น Kerry, Flash, ThaiPost';
COMMENT ON COLUMN orders.remark IS 'หมายเหตุเพิ่มเติมจากลูกค้า';
COMMENT ON COLUMN orders.ordered_at IS 'วันเวลาที่ทำรายการสั่งซื้อ';
