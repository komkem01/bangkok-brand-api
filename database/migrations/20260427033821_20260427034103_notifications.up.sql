SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE notifications (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	channel notification_channel NOT NULL DEFAULT 'in_app',
	status notification_status NOT NULL DEFAULT 'pending',
	title varchar NOT NULL,
	body text,
	image_url varchar,
	action_url varchar,
	ref_type varchar,
	ref_id uuid,
	read_at timestamp,
	sent_at timestamp,
	failed_reason varchar,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_notifications_member_id ON notifications(member_id);
CREATE INDEX IF NOT EXISTS idx_notifications_member_unread ON notifications(member_id, status) WHERE status != 'read';

--bun:split

COMMENT ON TABLE notifications IS 'ตารางบันทึกการแจ้งเตือนทุกช่องทาง (push/email/sms/in-app)';
COMMENT ON COLUMN notifications.member_id IS 'สมาชิกผู้รับการแจ้งเตือน';
COMMENT ON COLUMN notifications.channel IS 'ช่องทางที่ใช้ส่ง';
COMMENT ON COLUMN notifications.status IS 'สถานะการส่งและอ่าน';
COMMENT ON COLUMN notifications.title IS 'หัวข้อการแจ้งเตือน';
COMMENT ON COLUMN notifications.body IS 'เนื้อหาการแจ้งเตือน';
COMMENT ON COLUMN notifications.action_url IS 'Deep link หรือ URL ที่เปิดเมื่อกด';
COMMENT ON COLUMN notifications.ref_type IS 'ประเภทอ้างอิง เช่น order, return_request, dispute_case';
COMMENT ON COLUMN notifications.ref_id IS 'ไอดีของระเบียนที่อ้างอิง';
