SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE member_devices (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	platform varchar NOT NULL,
	device_token varchar NOT NULL,
	device_name varchar,
	app_version varchar,
	is_active boolean NOT NULL DEFAULT true,
	last_seen_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (member_id, device_token)
);

CREATE INDEX IF NOT EXISTS idx_member_devices_member_id ON member_devices(member_id);

--bun:split

COMMENT ON TABLE member_devices IS 'ตารางเก็บ device token สำหรับส่ง push notification';
COMMENT ON COLUMN member_devices.platform IS 'ระบบปฏิบัติการ เช่น ios, android, web';
COMMENT ON COLUMN member_devices.device_token IS 'FCM/APNs token สำหรับส่ง push notification';
COMMENT ON COLUMN member_devices.device_name IS 'ชื่ออุปกรณ์หรือโมเดลของอุปกรณ์';
COMMENT ON COLUMN member_devices.is_active IS 'อุปกรณ์นี้ยังเปิดใช้งานและรับการแจ้งเตือนอยู่หรือไม่';
COMMENT ON COLUMN member_devices.last_seen_at IS 'เวลาบันทึกสุดท้ายที่ token นี้ถูกใช้เช็คอิน';
