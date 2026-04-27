SET statement_timeout = 0;

--bun:split

ALTER TABLE cart_items
	ADD COLUMN IF NOT EXISTS variant_id uuid REFERENCES product_variants(id);

COMMENT ON COLUMN cart_items.variant_id IS 'ตัวเลือกย่อยของสินค้าที่ถูกเพิ่มลงตะกร้า';
CREATE INDEX IF NOT EXISTS idx_cart_items_variant_id ON cart_items(variant_id);

--bun:split

ALTER TABLE order_items
	ADD COLUMN IF NOT EXISTS variant_id uuid REFERENCES product_variants(id);

COMMENT ON COLUMN order_items.variant_id IS 'ตัวเลือกย่อยของสินค้าที่ถูกซื้อจริง';
CREATE INDEX IF NOT EXISTS idx_order_items_variant_id ON order_items(variant_id);

--bun:split

ALTER TABLE product_images
	ADD COLUMN IF NOT EXISTS variant_id uuid REFERENCES product_variants(id);

COMMENT ON COLUMN product_images.variant_id IS 'รูปภาพเฉพาะตัวเลือกย่อย (ถ้ามี)';
CREATE INDEX IF NOT EXISTS idx_product_images_variant_id ON product_images(variant_id);
