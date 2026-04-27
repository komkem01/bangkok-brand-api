SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE invoice_items (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	invoice_id uuid REFERENCES invoices(id),
	order_item_id uuid REFERENCES order_items(id),
	description varchar,
	quantity integer NOT NULL DEFAULT 1,
	unit_price decimal(12, 2) NOT NULL DEFAULT 0,
	discount_amount decimal(12, 2) NOT NULL DEFAULT 0,
	tax_amount decimal(12, 2) NOT NULL DEFAULT 0,
	total_amount decimal(12, 2) NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_invoice_items_invoice_id ON invoice_items(invoice_id);

--bun:split

COMMENT ON TABLE invoice_items IS 'ตารางรายการสินค้า/บริการในใบกำกับภาษี';
COMMENT ON COLUMN invoice_items.invoice_id IS 'ใบกำกับภาษีที่รายการนี้สังกัด';
COMMENT ON COLUMN invoice_items.order_item_id IS 'รายการสินค้าในออเดอร์ที่อ้างอิง';
COMMENT ON COLUMN invoice_items.description IS 'คำอธิบายรายการในใบกำกับภาษี';
COMMENT ON COLUMN invoice_items.unit_price IS 'ราคาต่อหน่วย (ยังไม่รวมภาษี)';
COMMENT ON COLUMN invoice_items.tax_amount IS 'ภาษีวาย VAT ของรายการนี้';
