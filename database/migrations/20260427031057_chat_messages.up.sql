SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE chat_messages (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	room_id uuid REFERENCES chat_rooms(id),
	sender_id uuid REFERENCES members(id),
	type message_type NOT NULL DEFAULT 'text',
	message text,
	storage_id uuid REFERENCES storages(id),
	is_read boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE chat_messages IS 'ตารางเก็บข้อความสนทนาทั้งหมด';
COMMENT ON COLUMN chat_messages.message IS 'เนื้อหาข้อความ';
COMMENT ON COLUMN chat_messages.storage_id IS 'อ้างอิงไฟล์ภาพหรือเอกสาร (ถ้ามี)';
