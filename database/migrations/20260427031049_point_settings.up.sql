SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE point_settings (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	earn_rate_amount decimal(12, 2),
	earn_points integer,
	min_order_to_earn decimal(12, 2) NOT NULL DEFAULT 0,
	point_expiry_months integer,
	is_active boolean NOT NULL DEFAULT true,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE point_settings IS 'ตารางตั้งค่าเกณฑ์การให้คะแนนสะสม';
COMMENT ON COLUMN point_settings.earn_rate_amount IS 'ยอดเงินที่ทำให้ได้คะแนน (เช่น ทุก 25 บาท)';
COMMENT ON COLUMN point_settings.earn_points IS 'จำนวนคะแนนที่ได้รับ (เช่น ได้ 1 คะแนน)';
COMMENT ON COLUMN point_settings.min_order_to_earn IS 'ยอดขั้นต่ำที่จะเริ่มสะสมคะแนนได้';
COMMENT ON COLUMN point_settings.point_expiry_months IS 'อายุของคะแนน (เดือน)';
