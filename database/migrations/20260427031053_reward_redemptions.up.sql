SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE reward_redemptions (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	reward_id uuid REFERENCES rewards(id),
	points_used integer,
	status redemption_status NOT NULL DEFAULT 'pending',
	recipient_name varchar,
	recipient_phone varchar,
	shipping_address text,
	redeemed_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE reward_redemptions IS 'ตารางบันทึกประวัติการแลกของรางวัล';
COMMENT ON COLUMN reward_redemptions.points_used IS 'คะแนนที่ใช้แลกจริง';
COMMENT ON COLUMN reward_redemptions.status IS 'สถานะการแลกรางวัล';
COMMENT ON COLUMN reward_redemptions.recipient_name IS 'ชื่อผู้รับของรางวัล';
COMMENT ON COLUMN reward_redemptions.recipient_phone IS 'เบอร์โทรศัพท์ผู้รับของรางวัล';
COMMENT ON COLUMN reward_redemptions.shipping_address IS 'ที่อยู่จัดส่งของรางวัล';
