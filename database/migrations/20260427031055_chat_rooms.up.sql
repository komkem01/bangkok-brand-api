SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE chat_rooms (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	order_id uuid REFERENCES orders(id),
	brand_id uuid REFERENCES brands(id),
	last_message text,
	last_message_at timestamp,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE chat_rooms IS 'ตารางเก็บห้องสนทนาหลัก';
COMMENT ON COLUMN chat_rooms.order_id IS 'อ้างอิงคำสั่งซื้อ (ถ้าเป็นการคุยเรื่องออเดอร์)';
COMMENT ON COLUMN chat_rooms.brand_id IS 'อ้างอิงแบรนด์หรือร้านค้าที่ลูกค้าคุยด้วย';
COMMENT ON COLUMN chat_rooms.last_message IS 'ข้อความล่าสุดเพื่อแสดงในหน้า Preview';
