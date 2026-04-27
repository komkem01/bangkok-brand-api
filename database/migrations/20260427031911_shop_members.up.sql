SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shop_members (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid REFERENCES shops(id),
	member_id uuid REFERENCES members(id),
	role shop_member_role NOT NULL DEFAULT 'staff',
	is_active boolean NOT NULL DEFAULT true,
	joined_at timestamp,
	invited_by_id uuid REFERENCES members(id),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (shop_id, member_id)
);

--bun:split

COMMENT ON TABLE shop_members IS 'ตารางเก็บสมาชิกที่มีสิทธิ์จัดการร้านค้า';
COMMENT ON COLUMN shop_members.shop_id IS 'ร้านค้าที่สมาชิกสังกัด';
COMMENT ON COLUMN shop_members.member_id IS 'สมาชิกผู้ดูแลร้าน';
COMMENT ON COLUMN shop_members.role IS 'บทบาทของสมาชิกในร้าน';
COMMENT ON COLUMN shop_members.invited_by_id IS 'ผู้ที่เชิญสมาชิกเข้าร้าน';
