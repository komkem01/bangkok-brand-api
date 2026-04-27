SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE kyc_verifications (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid UNIQUE REFERENCES members(id),
	entity_type kyc_entity_type NOT NULL DEFAULT 'merchant',
	legal_name varchar,
	business_name varchar,
	citizen_or_tax_id varchar,
	contact_phone varchar,
	contact_email varchar,
	status kyc_status NOT NULL DEFAULT 'pending',
	submitted_at timestamp,
	reviewed_at timestamp,
	reviewer_id uuid REFERENCES members(id),
	rejection_reason text,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE kyc_verifications IS 'ตารางหลักสำหรับตรวจสอบและยืนยันตัวตนของร้านค้าหรือกลุ่มวิสาหกิจชุมชน';
COMMENT ON COLUMN kyc_verifications.member_id IS 'สมาชิกเจ้าของข้อมูล KYC';
COMMENT ON COLUMN kyc_verifications.entity_type IS 'ประเภทผู้ยื่น KYC';
COMMENT ON COLUMN kyc_verifications.legal_name IS 'ชื่อบุคคลหรือนิติบุคคลตามเอกสาร';
COMMENT ON COLUMN kyc_verifications.business_name IS 'ชื่อร้านค้าหรือกลุ่มวิสาหกิจ';
COMMENT ON COLUMN kyc_verifications.citizen_or_tax_id IS 'เลขบัตรประชาชนหรือเลขประจำตัวผู้เสียภาษี';
COMMENT ON COLUMN kyc_verifications.status IS 'สถานะปัจจุบันของการตรวจสอบ KYC';
COMMENT ON COLUMN kyc_verifications.reviewer_id IS 'ผู้ตรวจสอบหรืออนุมัติ KYC';
COMMENT ON COLUMN kyc_verifications.rejection_reason IS 'เหตุผลการไม่อนุมัติหรือขอเอกสารเพิ่ม';
