SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE member_bank_accounts (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	bank_id uuid REFERENCES banks(id),
	account_number varchar,
	account_name varchar,
	branch_name varchar,
	is_default boolean NOT NULL DEFAULT false,
	is_verified boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE member_bank_accounts IS 'ตารางเก็บข้อมูลบัญชีธนาคารของสมาชิก';
COMMENT ON COLUMN member_bank_accounts.member_id IS 'อ้างอิงสมาชิกเจ้าของบัญชี';
COMMENT ON COLUMN member_bank_accounts.bank_id IS 'อ้างอิงธนาคารของบัญชี';
COMMENT ON COLUMN member_bank_accounts.account_number IS 'เลขที่บัญชีธนาคาร';
COMMENT ON COLUMN member_bank_accounts.account_name IS 'ชื่อบัญชีธนาคาร';
COMMENT ON COLUMN member_bank_accounts.branch_name IS 'ชื่อสาขาธนาคาร';
COMMENT ON COLUMN member_bank_accounts.is_default IS 'ระบุว่าเป็นบัญชีหลักของสมาชิก';
COMMENT ON COLUMN member_bank_accounts.is_verified IS 'สถานะการยืนยันบัญชีธนาคาร';
