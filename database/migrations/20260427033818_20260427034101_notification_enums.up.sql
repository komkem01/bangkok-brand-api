SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'notification_channel') THEN
		CREATE TYPE notification_channel AS ENUM (
			'push',
			'email',
			'sms',
			'in_app'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'notification_status') THEN
		CREATE TYPE notification_status AS ENUM (
			'pending',
			'sent',
			'delivered',
			'read',
			'failed'
		);
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE notification_channel IS 'ช่องทางการแจ้งเตือน push/email/sms/in-app';
COMMENT ON TYPE notification_status IS 'สถานะการส่งและอ่านการแจ้งเตือน';
