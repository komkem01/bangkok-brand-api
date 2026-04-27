SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE kyc_documents (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	kyc_verification_id uuid REFERENCES kyc_verifications(id),
	document_type kyc_document_type,
	storage_id uuid REFERENCES storages(id),
	document_no varchar,
	issued_at timestamp,
	expired_at timestamp,
	is_verified boolean NOT NULL DEFAULT false,
	verified_at timestamp,
	note text,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE kyc_documents IS 'ตารางเก็บเอกสารประกอบการยืนยันตัวตน KYC';
COMMENT ON COLUMN kyc_documents.kyc_verification_id IS 'อ้างอิงคำขอ KYC';
COMMENT ON COLUMN kyc_documents.document_type IS 'ประเภทเอกสาร KYC';
COMMENT ON COLUMN kyc_documents.storage_id IS 'ไฟล์เอกสารที่เก็บในระบบ storage';
COMMENT ON COLUMN kyc_documents.document_no IS 'เลขเอกสาร';
COMMENT ON COLUMN kyc_documents.is_verified IS 'สถานะการตรวจสอบเอกสารรายไฟล์';
COMMENT ON COLUMN kyc_documents.note IS 'หมายเหตุเพิ่มเติมของเอกสาร';
