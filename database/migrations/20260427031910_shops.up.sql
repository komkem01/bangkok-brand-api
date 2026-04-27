SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shops (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	owner_member_id uuid REFERENCES members(id),
	brand_id uuid UNIQUE REFERENCES brands(id),
	kyc_verification_id uuid UNIQUE REFERENCES kyc_verifications(id),
	shop_code varchar UNIQUE,
	name_th varchar,
	name_en varchar,
	slug varchar UNIQUE,
	description text,
	logo_id uuid REFERENCES storages(id),
	cover_image_id uuid REFERENCES storages(id),
	contact_phone varchar,
	contact_email varchar,
	address_detail text,
	province_id uuid REFERENCES provinces(id),
	district_id uuid REFERENCES districts(id),
	sub_district_id uuid REFERENCES sub_districts(id),
	zipcode_id uuid REFERENCES zipcodes(id),
	status shop_status NOT NULL DEFAULT 'pending_kyc',
	is_active boolean NOT NULL DEFAULT true,
	opened_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE shops IS 'ตารางหลักของร้านค้า/ผู้ขายในระบบ B2C';
COMMENT ON COLUMN shops.owner_member_id IS 'เจ้าของร้านค้าหลัก';
COMMENT ON COLUMN shops.brand_id IS 'อ้างอิงแบรนด์ประจำร้าน';
COMMENT ON COLUMN shops.kyc_verification_id IS 'อ้างอิงผลการยืนยันตัวตน KYC';
COMMENT ON COLUMN shops.shop_code IS 'รหัสร้านค้า';
COMMENT ON COLUMN shops.slug IS 'ชื่อย่อร้านสำหรับ URL';
COMMENT ON COLUMN shops.logo_id IS 'โลโก้ร้านค้า';
COMMENT ON COLUMN shops.cover_image_id IS 'ภาพหน้าปกร้านค้า';
COMMENT ON COLUMN shops.status IS 'สถานะการดำเนินงานของร้าน';
