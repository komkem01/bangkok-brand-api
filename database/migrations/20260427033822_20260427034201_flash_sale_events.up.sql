SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'flash_sale_status') THEN
		CREATE TYPE flash_sale_status AS ENUM (
			'scheduled',
			'active',
			'ended',
			'cancelled'
		);
	END IF;
END
$$;

COMMENT ON TYPE flash_sale_status IS 'สถานะเอเวนต์ flash sale';

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE flash_sale_events (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid REFERENCES shops(id),
	name varchar NOT NULL,
	description text,
	cover_storage_id uuid REFERENCES storages(id),
	status flash_sale_status NOT NULL DEFAULT 'scheduled',
	starts_at timestamp NOT NULL,
	ends_at timestamp NOT NULL,
	max_orders_per_member integer,
	is_visible boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_flash_sale_events_shop_id ON flash_sale_events(shop_id);
CREATE INDEX IF NOT EXISTS idx_flash_sale_events_status ON flash_sale_events(status);
CREATE INDEX IF NOT EXISTS idx_flash_sale_events_active_window ON flash_sale_events(starts_at, ends_at);

--bun:split

COMMENT ON TABLE flash_sale_events IS 'ตารางเอเวนต์ flash sale สินค้าลดราคาช่วงเวลาจำกัด';
COMMENT ON COLUMN flash_sale_events.shop_id IS 'ร้านค้าสังกัดเอเวนต์ (null = platform-wide)';
COMMENT ON COLUMN flash_sale_events.name IS 'ชื่อเอเวนต์';
COMMENT ON COLUMN flash_sale_events.starts_at IS 'เวลาเริ่มต้น flash sale';
COMMENT ON COLUMN flash_sale_events.ends_at IS 'เวลาสิ้นสุด flash sale';
COMMENT ON COLUMN flash_sale_events.max_orders_per_member IS 'จำนวนสิทธิซื้อสูงสุดต่อออเดอร์ต่อสมาชิก';
