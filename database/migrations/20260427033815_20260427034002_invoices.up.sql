SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE invoices (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	invoice_no varchar UNIQUE,
	order_id uuid REFERENCES orders(id),
	payment_id uuid REFERENCES payments(id),
	member_id uuid REFERENCES members(id),
	shop_id uuid REFERENCES shops(id),
	type invoice_type NOT NULL DEFAULT 'receipt',
	status invoice_status NOT NULL DEFAULT 'draft',
	issue_date date,
	due_date date,
	sub_total decimal(12, 2) NOT NULL DEFAULT 0,
	discount_amount decimal(12, 2) NOT NULL DEFAULT 0,
	tax_amount decimal(12, 2) NOT NULL DEFAULT 0,
	total_amount decimal(12, 2) NOT NULL DEFAULT 0,
	tax_rate decimal(5, 4) NOT NULL DEFAULT 0.07,
	billing_name varchar,
	billing_address text,
	billing_tax_id varchar,
	note text,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_invoices_order_id ON invoices(order_id);
CREATE INDEX IF NOT EXISTS idx_invoices_member_id ON invoices(member_id);
CREATE INDEX IF NOT EXISTS idx_invoices_shop_id ON invoices(shop_id);

--bun:split

COMMENT ON TABLE invoices IS 'ตารางใบกำกับภาษี/ใบเสร็จสำหรับออเดอร์แต่ละรายการ';
COMMENT ON COLUMN invoices.invoice_no IS 'เลขที่ใบกำกับภาษี/ใบเสร็จเรียงตามลำดับ';
COMMENT ON COLUMN invoices.order_id IS 'ออเดอร์ที่อ้างอิง';
COMMENT ON COLUMN invoices.payment_id IS 'รายการชำระเงินที่เกี่ยวข้อง';
COMMENT ON COLUMN invoices.type IS 'ประเภทเอกสาร เช่น ใบกำกับภาษีเต็มรูป, ใบเสร็จรับเงิน';
COMMENT ON COLUMN invoices.status IS 'สถานะของเอกสาร';
COMMENT ON COLUMN invoices.sub_total IS 'ยอดก่อนส่วนลดและภาษี';
COMMENT ON COLUMN invoices.tax_amount IS 'ยอดภาษีมูลค่าเพิ่ม (VAT)';
COMMENT ON COLUMN invoices.tax_rate IS 'อัตราภาษี เช่น 0.07 = 7%';
COMMENT ON COLUMN invoices.billing_name IS 'ชื่อองค์กร/บุคคลสำหรับออกใบกำกับภาษี';
COMMENT ON COLUMN invoices.billing_tax_id IS 'เลขประจำตัวผู้เสียภาษีสำหรับออกใบกำกับภาษี';
