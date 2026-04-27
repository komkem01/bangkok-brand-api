SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'audit_action') THEN
		CREATE TYPE audit_action AS ENUM ('INSERT', 'UPDATE', 'DELETE');
	END IF;
END
$$;

COMMENT ON TYPE audit_action IS 'ประเภทการเปลี่ยนแปลงข้อมูล';

--bun:split

CREATE TABLE audit_logs (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	table_name varchar NOT NULL,
	record_id uuid,
	action audit_action NOT NULL,
	actor_id uuid,
	actor_type varchar,
	old_values jsonb,
	new_values jsonb,
	changed_fields text[],
	ip_address inet,
	user_agent varchar,
	request_id varchar,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_audit_logs_table_record ON audit_logs(table_name, record_id);
CREATE INDEX IF NOT EXISTS idx_audit_logs_actor_id ON audit_logs(actor_id);
CREATE INDEX IF NOT EXISTS idx_audit_logs_created_at ON audit_logs(created_at);

--bun:split

COMMENT ON TABLE audit_logs IS 'Append-only audit log บันทึกการเปลี่ยนแปลงข้อมูลในตาราง sensitive';
COMMENT ON COLUMN audit_logs.table_name IS 'ชื่อตารางที่เกิดการเปลี่ยนแปลง';
COMMENT ON COLUMN audit_logs.record_id IS 'Primary key ของ row ที่เปลี่ยนแปลง';
COMMENT ON COLUMN audit_logs.action IS 'ประเภท operation: INSERT / UPDATE / DELETE';
COMMENT ON COLUMN audit_logs.actor_id IS 'ID ของผู้ทำรายการ (member/admin)';
COMMENT ON COLUMN audit_logs.actor_type IS 'ประเภทผู้กระทำ เช่น member, admin, system, service';
COMMENT ON COLUMN audit_logs.old_values IS 'ค่าก่อนการเปลี่ยนแปลง (UPDATE/DELETE)';
COMMENT ON COLUMN audit_logs.new_values IS 'ค่าหลังการเปลี่ยนแปลง (INSERT/UPDATE)';
COMMENT ON COLUMN audit_logs.changed_fields IS 'รายชื่อ column ที่ถูกเปลี่ยนแปลงใน UPDATE';
COMMENT ON COLUMN audit_logs.ip_address IS 'IP address ของผู้ทำรายการ';
COMMENT ON COLUMN audit_logs.request_id IS 'Trace/Request ID สำหรับเชื่อมโยงกับ application logs';
