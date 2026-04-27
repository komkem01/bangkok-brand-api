SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shop_settings (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shop_id uuid UNIQUE REFERENCES shops(id),
	auto_accept_orders boolean NOT NULL DEFAULT false,
	allow_cod boolean NOT NULL DEFAULT true,
	min_order_amount decimal(12, 2) NOT NULL DEFAULT 0,
	preparation_time_minutes integer NOT NULL DEFAULT 30,
	default_shipping_method_id uuid REFERENCES shipping_methods(id),
	return_policy text,
	refund_policy text,
	business_hours json,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE shop_settings IS 'ตารางตั้งค่าการดำเนินงานของแต่ละร้านค้า';
COMMENT ON COLUMN shop_settings.auto_accept_orders IS 'ร้านยืนยันออเดอร์อัตโนมัติหรือไม่';
COMMENT ON COLUMN shop_settings.allow_cod IS 'ร้านรองรับเก็บเงินปลายทางหรือไม่';
COMMENT ON COLUMN shop_settings.min_order_amount IS 'ยอดสั่งซื้อขั้นต่ำของร้าน';
COMMENT ON COLUMN shop_settings.preparation_time_minutes IS 'เวลาเตรียมสินค้าโดยประมาณ (นาที)';
COMMENT ON COLUMN shop_settings.default_shipping_method_id IS 'ช่องทางจัดส่งเริ่มต้นของร้าน';
COMMENT ON COLUMN shop_settings.business_hours IS 'เวลาทำการของร้านในรูปแบบ JSON';
