SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE products (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	category_id uuid REFERENCES categories(id),
	brand_id uuid REFERENCES brands(id),
	merchant_id uuid,
	sku varchar UNIQUE,
	name_th varchar,
	name_en varchar,
	short_description_th text,
	full_description_th text,
	price decimal(12, 2),
	discount_price decimal(12, 2),
	is_on_sale boolean NOT NULL DEFAULT false,
	slug varchar UNIQUE,
	meta_title varchar,
	meta_description text,
	status varchar,
	is_active boolean NOT NULL DEFAULT true,
	is_featured boolean NOT NULL DEFAULT false,
	weight decimal(8, 2),
	width decimal(8, 2),
	length decimal(8, 2),
	height decimal(8, 2),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	deleted_at timestamp
);

--bun:split

COMMENT ON TABLE products IS 'ตารางหลักเก็บข้อมูลคุณสมบัติของสินค้าทั้งหมด';
COMMENT ON COLUMN products.category_id IS 'หมวดหมู่ที่สินค้าสังกัด';
COMMENT ON COLUMN products.brand_id IS 'แบรนด์ของสินค้า';
COMMENT ON COLUMN products.merchant_id IS 'เจ้าของร้านค้าที่ลงขายสินค้า (รอเชื่อมกับตาราง Shops)';
COMMENT ON COLUMN products.sku IS 'รหัสควบคุมสต็อก (Stock Keeping Unit)';
COMMENT ON COLUMN products.name_th IS 'ชื่อสินค้าภาษาไทย';
COMMENT ON COLUMN products.name_en IS 'ชื่อสินค้าภาษาอังกฤษ';
COMMENT ON COLUMN products.short_description_th IS 'รายละเอียดสินค้าแบบย่อ';
COMMENT ON COLUMN products.full_description_th IS 'รายละเอียดสินค้าแบบเต็ม';
COMMENT ON COLUMN products.price IS 'ราคาขายปกติ';
COMMENT ON COLUMN products.discount_price IS 'ราคาหลังหักส่วนลด (ถ้ามี)';
COMMENT ON COLUMN products.is_on_sale IS 'สถานะการจัดโปรโมชั่นลดราคา';
COMMENT ON COLUMN products.slug IS 'Path URL สำหรับเข้าถึงหน้านั้นๆ ของสินค้า';
COMMENT ON COLUMN products.meta_title IS 'ชื่อหัวข้อสำหรับการทำ SEO';
COMMENT ON COLUMN products.meta_description IS 'รายละเอียดคำอธิบายสำหรับการทำ SEO';
COMMENT ON COLUMN products.status IS 'สถานะการขาย (เช่น active, draft, archived)';
COMMENT ON COLUMN products.is_active IS 'เปิด/ปิด การมองเห็นสินค้า';
COMMENT ON COLUMN products.is_featured IS 'แสดงเป็นสินค้าแนะนำในหน้าหลัก';
COMMENT ON COLUMN products.weight IS 'น้ำหนักสินค้า (หน่วยเป็นกิโลกรัม)';
COMMENT ON COLUMN products.width IS 'ความกว้างของบรรจุภัณฑ์ (ซม.)';
COMMENT ON COLUMN products.length IS 'ความยาวของบรรจุภัณฑ์ (ซม.)';
COMMENT ON COLUMN products.height IS 'ความสูงของบรรจุภัณฑ์ (ซม.)';
