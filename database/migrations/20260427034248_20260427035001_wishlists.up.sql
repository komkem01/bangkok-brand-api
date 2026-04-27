SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE wishlists (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid NOT NULL REFERENCES members(id),
	product_id uuid NOT NULL REFERENCES products(id),
	product_variant_id uuid,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (member_id, product_id, product_variant_id)
);

CREATE INDEX IF NOT EXISTS idx_wishlists_member_id ON wishlists(member_id);
CREATE INDEX IF NOT EXISTS idx_wishlists_product_id ON wishlists(product_id);

--bun:split

COMMENT ON TABLE wishlists IS 'ตารางรายการสินค้าที่ลูกค้าบันทึก/กดหัวใจเอาไว้';
COMMENT ON COLUMN wishlists.member_id IS 'สมาชิกที่บันทึกสินค้า';
COMMENT ON COLUMN wishlists.product_id IS 'สินค้าที่ถูกบันทึก';
COMMENT ON COLUMN wishlists.product_variant_id IS 'ตัวเลือก variant ที่สนใจ (null = สนใจสินค้าทั้งหมด)';
