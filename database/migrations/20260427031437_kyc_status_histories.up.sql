SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE kyc_status_histories (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	kyc_verification_id uuid REFERENCES kyc_verifications(id),
	old_status kyc_status,
	new_status kyc_status,
	changed_by_id uuid REFERENCES members(id),
	remark text,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE kyc_status_histories IS 'ตารางบันทึกประวัติการเปลี่ยนสถานะ KYC';
COMMENT ON COLUMN kyc_status_histories.old_status IS 'สถานะเดิมก่อนการเปลี่ยนแปลง';
COMMENT ON COLUMN kyc_status_histories.new_status IS 'สถานะใหม่หลังการเปลี่ยนแปลง';
COMMENT ON COLUMN kyc_status_histories.changed_by_id IS 'ผู้ที่เปลี่ยนสถานะ KYC';
COMMENT ON COLUMN kyc_status_histories.remark IS 'หมายเหตุการเปลี่ยนสถานะ';
