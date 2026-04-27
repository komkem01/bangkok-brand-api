SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE dispute_messages (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	dispute_case_id uuid REFERENCES dispute_cases(id),
	sender_member_id uuid REFERENCES members(id),
	party dispute_message_party,
	message text,
	evidence_storage_id uuid REFERENCES storages(id),
	is_internal_note boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_dispute_messages_case_id ON dispute_messages(dispute_case_id);

--bun:split

COMMENT ON TABLE dispute_messages IS 'ตารางข้อความและบันทึกในเคสข้อพิพาท';
COMMENT ON COLUMN dispute_messages.dispute_case_id IS 'เคสข้อพิพาทที่ข้อความนี้สังกัด';
COMMENT ON COLUMN dispute_messages.sender_member_id IS 'สมาชิกผู้ส่งข้อความ';
COMMENT ON COLUMN dispute_messages.party IS 'บทบาทฝั่งผู้ส่งในเคส';
COMMENT ON COLUMN dispute_messages.message IS 'เนื้อหาข้อความ';
COMMENT ON COLUMN dispute_messages.evidence_storage_id IS 'หลักฐานแนบในข้อความ';
COMMENT ON COLUMN dispute_messages.is_internal_note IS 'หมายเหตุภายในสำหรับแอดมินเท่านั้น';
