SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE idempotency_keys (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	key varchar NOT NULL UNIQUE,
	method varchar NOT NULL,
	path varchar NOT NULL,
	request_hash varchar NOT NULL,
	status_code integer,
	response_body jsonb,
	processed_at timestamp,
	expires_at timestamp NOT NULL,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_idempotency_keys_expires_at ON idempotency_keys(expires_at);

--bun:split

COMMENT ON TABLE idempotency_keys IS 'ตารางเก็บ idempotency key สำหรับป้องกันการดำเนินการซ้ำซ้อน (duplicate requests)';
COMMENT ON COLUMN idempotency_keys.key IS 'Idempotency key ที่ client ส่งมาใน header X-Idempotency-Key';
COMMENT ON COLUMN idempotency_keys.method IS 'HTTP method เช่น POST, PUT';
COMMENT ON COLUMN idempotency_keys.path IS 'Request path ที่ถูกเรียก';
COMMENT ON COLUMN idempotency_keys.request_hash IS 'Hash ของ request body เพื่อตรวจสอบความเหมือนกัน';
COMMENT ON COLUMN idempotency_keys.status_code IS 'HTTP status code ที่ตอบกลับครั้งแรก';
COMMENT ON COLUMN idempotency_keys.response_body IS 'Response body ที่ตอบกลับครั้งแรก เพื่อนำกลับมาตอบซ้ำได้เลย';
COMMENT ON COLUMN idempotency_keys.expires_at IS 'วันเวลาหมดอายุของ key เพื่อสามารถลบบันทึกเก่าได้';
