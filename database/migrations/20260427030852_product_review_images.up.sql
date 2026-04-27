SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE product_review_images (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	review_id uuid REFERENCES product_reviews(id),
	storage_id uuid REFERENCES storages(id),
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE product_review_images IS 'ตารางเก็บรูปภาพที่ลูกค้าอัปโหลดพร้อมรีวิว';
COMMENT ON COLUMN product_review_images.review_id IS 'อ้างอิงรีวิวเจ้าของรูปภาพ';
COMMENT ON COLUMN product_review_images.storage_id IS 'รูปภาพประกอบรีวิว';
