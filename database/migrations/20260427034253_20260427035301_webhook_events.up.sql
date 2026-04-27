SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'webhook_event_status') THEN
		CREATE TYPE webhook_event_status AS ENUM (
			'pending',
			'processing',
			'delivered',
			'failed',
			'abandoned'
		);
	END IF;
END
$$;

COMMENT ON TYPE webhook_event_status IS 'สถานะการส่ง webhook event';

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE webhook_events (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	event_type varchar NOT NULL,
	payload jsonb NOT NULL,
	status webhook_event_status NOT NULL DEFAULT 'pending',
	shop_id uuid REFERENCES shops(id),
	endpoint_url varchar NOT NULL,
	secret_hash varchar,
	attempt_count integer NOT NULL DEFAULT 0,
	max_attempts integer NOT NULL DEFAULT 5,
	last_attempt_at timestamp,
	next_retry_at timestamp,
	last_response_code integer,
	last_response_body text,
	delivered_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_webhook_events_shop_id ON webhook_events(shop_id);
CREATE INDEX IF NOT EXISTS idx_webhook_events_status ON webhook_events(status);
CREATE INDEX IF NOT EXISTS idx_webhook_events_pending_retry ON webhook_events(next_retry_at) WHERE status IN ('pending', 'failed');

--bun:split

COMMENT ON TABLE webhook_events IS 'ตาราง webhook ออกแจ้งเหตุการณ์ไปยังระบบภายนอก (outbound webhook)';
COMMENT ON COLUMN webhook_events.event_type IS 'ชื่อ event เช่น order.created, payment.completed, return.approved';
COMMENT ON COLUMN webhook_events.payload IS 'เนื้อหา event ที่ส่งไปในรูป JSON';
COMMENT ON COLUMN webhook_events.shop_id IS 'ร้านค้าเจ้าของ webhook (null = platform-level)';
COMMENT ON COLUMN webhook_events.endpoint_url IS 'URL ปลายทางที่จะส่ง POST';
COMMENT ON COLUMN webhook_events.secret_hash IS 'HMAC-SHA256 signature header สำหรับผู้รับตรวจสอบความถูกต้อง';
COMMENT ON COLUMN webhook_events.attempt_count IS 'จำนวนครั้งที่พยายามส่งแล้ว';
COMMENT ON COLUMN webhook_events.max_attempts IS 'จำนวนครั้งสูงสุดที่อนุญาต หลังจากนั้น status = abandoned';
COMMENT ON COLUMN webhook_events.next_retry_at IS 'เวลานัดส่งซ้ำครั้งถัดไป (exponential backoff)';
COMMENT ON COLUMN webhook_events.last_response_code IS 'HTTP response code ที่ได้รับจากปลายทางล่าสุด';
