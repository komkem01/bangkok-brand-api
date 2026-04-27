SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE admin_action_logs (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	admin_member_id uuid REFERENCES members(id),
	action_type admin_action_type,
	resource_type varchar,
	resource_id uuid,
	shop_id uuid REFERENCES shops(id),
	target_member_id uuid REFERENCES members(id),
	before_data json,
	after_data json,
	ip_address varchar,
	user_agent text,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE admin_action_logs IS 'ตารางบันทึกการกระทำของแอดมินเพื่อการตรวจสอบย้อนหลัง';
COMMENT ON COLUMN admin_action_logs.admin_member_id IS 'แอดมินผู้ดำเนินการ';
COMMENT ON COLUMN admin_action_logs.action_type IS 'ประเภทกิจกรรมที่ดำเนินการ';
COMMENT ON COLUMN admin_action_logs.resource_type IS 'ประเภททรัพยากรที่ถูกแก้ไข';
COMMENT ON COLUMN admin_action_logs.resource_id IS 'รหัสทรัพยากรที่ถูกแก้ไข';
COMMENT ON COLUMN admin_action_logs.before_data IS 'ข้อมูลก่อนแก้ไข';
COMMENT ON COLUMN admin_action_logs.after_data IS 'ข้อมูลหลังแก้ไข';
