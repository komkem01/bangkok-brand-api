SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_reviews (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	product_id uuid REFERENCES products(id),
	member_id uuid REFERENCES members(id),
	order_id uuid REFERENCES orders(id),
	rating integer,
	comment text,
	is_verified_purchase boolean NOT NULL DEFAULT true,
	is_hidden boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_reviews IS 'ตารางเก็บคะแนนและความคิดเห็นจากลูกค้า';
COMMENT ON COLUMN product_reviews.product_id IS 'สินค้าที่ถูกรีวิว';
COMMENT ON COLUMN product_reviews.member_id IS 'ผู้รีวิว';
COMMENT ON COLUMN product_reviews.order_id IS 'อ้างอิงออเดอร์ (ต้องซื้อจริงถึงรีวิวได้)';
COMMENT ON COLUMN product_reviews.rating IS 'คะแนน 1-5 ดาว';
COMMENT ON COLUMN product_reviews.comment IS 'ข้อความรีวิว';
COMMENT ON COLUMN product_reviews.is_verified_purchase IS 'รีวิวนี้มาจากการซื้อสินค้าจริง';
COMMENT ON COLUMN product_reviews.is_hidden IS 'Admin สามารถสั่งซ่อนรีวิวที่ไม่เหมาะสมได้';
