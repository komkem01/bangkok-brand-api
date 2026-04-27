SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_views (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	session_id varchar,
	product_id uuid NOT NULL REFERENCES products(id),
	product_variant_id uuid,
	ref_source varchar,
	platform varchar,
	viewed_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_product_views_member_id ON product_views(member_id);
CREATE INDEX IF NOT EXISTS idx_product_views_product_id ON product_views(product_id);
CREATE INDEX IF NOT EXISTS idx_product_views_viewed_at ON product_views(viewed_at);

--bun:split

COMMENT ON TABLE product_views IS 'ตารางบันทึกการเข้าชมหน้าสินค้า ใช้สำหรับ recommendation engine';
COMMENT ON COLUMN product_views.member_id IS 'สมาชิกที่เข้าชม (null = anonymous)';
COMMENT ON COLUMN product_views.session_id IS 'Anonymous session ID สำหรับผู้ใช้ทั่วไป';
COMMENT ON COLUMN product_views.product_id IS 'สินค้าที่ถูกเข้าชม';
COMMENT ON COLUMN product_views.ref_source IS 'แหล่งที่มาของการคลิก เช่น search, homepage_banner, flash_sale, recommendation';
