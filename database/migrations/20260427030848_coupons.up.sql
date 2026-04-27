SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE coupons (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	code varchar UNIQUE,
	name_th varchar,
	description text,
	type coupon_type,
	value decimal(12, 2),
	max_discount_amount decimal(12, 2),
	min_order_amount decimal(12, 2) NOT NULL DEFAULT 0,
	limit_per_coupon integer,
	limit_per_member integer NOT NULL DEFAULT 1,
	used_count integer NOT NULL DEFAULT 0,
	start_date timestamp,
	end_date timestamp,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE coupons IS 'ตารางเก็บรหัสส่วนลดและเงื่อนไขโปรโมชั่น';
COMMENT ON COLUMN coupons.code IS 'รหัสคูปองที่ลูกค้าต้องกรอก (เช่น BKK2024)';
COMMENT ON COLUMN coupons.name_th IS 'ชื่อแคมเปญคูปอง';
COMMENT ON COLUMN coupons.description IS 'เงื่อนไขการใช้คูปอง';
COMMENT ON COLUMN coupons.type IS 'ประเภทส่วนลด';
COMMENT ON COLUMN coupons.value IS 'มูลค่าส่วนลด (เงินหรือเปอร์เซ็นต์)';
COMMENT ON COLUMN coupons.max_discount_amount IS 'ส่วนลดสูงสุด (กรณีเป็นเปอร์เซ็นต์)';
COMMENT ON COLUMN coupons.min_order_amount IS 'ยอดสั่งซื้อขั้นต่ำที่ใช้คูปองได้';
COMMENT ON COLUMN coupons.limit_per_coupon IS 'จำนวนครั้งที่คูปองนี้ใช้ได้ทั้งหมด (ถ้ามี)';
COMMENT ON COLUMN coupons.limit_per_member IS 'จำนวนครั้งที่สมาชิกแต่ละคนใช้ได้';
COMMENT ON COLUMN coupons.used_count IS 'จำนวนครั้งที่ถูกใช้ไปแล้ว';
COMMENT ON COLUMN coupons.start_date IS 'วันที่เริ่มใช้งานได้';
COMMENT ON COLUMN coupons.end_date IS 'วันที่คูปองหมดอายุ';
