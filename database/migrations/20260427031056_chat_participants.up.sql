SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE chat_participants (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	room_id uuid REFERENCES chat_rooms(id),
	member_id uuid REFERENCES members(id),
	joined_at timestamp,
	last_read_at timestamp,
	is_active boolean NOT NULL DEFAULT true
);

--bun:split

COMMENT ON TABLE chat_participants IS 'ตารางเก็บผู้เข้าร่วมในแต่ละห้องแชท';
COMMENT ON COLUMN chat_participants.last_read_at IS 'วันเวลาที่อ่านล่าสุด เพื่อคำนวณ Unread Count';
