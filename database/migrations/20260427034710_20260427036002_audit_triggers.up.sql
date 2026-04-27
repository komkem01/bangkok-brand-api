SET statement_timeout = 0;

--bun:split

-- ฟังก์ชัน trigger สำหรับบันทึก audit log อัตโนมัติ
CREATE OR REPLACE FUNCTION fn_audit_log()
RETURNS TRIGGER AS $$
DECLARE
	v_old_values jsonb := NULL;
	v_new_values jsonb := NULL;
	v_changed_fields text[] := NULL;
	v_record_id uuid;
BEGIN
	IF TG_OP = 'INSERT' THEN
		v_new_values := to_jsonb(NEW);
		v_record_id := (NEW).id;
	ELSIF TG_OP = 'UPDATE' THEN
		v_old_values := to_jsonb(OLD);
		v_new_values := to_jsonb(NEW);
		v_record_id := (NEW).id;
		-- เก็บเฉพาะ field ที่เปลี่ยนแปลง
		SELECT array_agg(key)
		INTO v_changed_fields
		FROM (
			SELECT key
			FROM jsonb_each(to_jsonb(NEW))
			WHERE to_jsonb(NEW) -> key IS DISTINCT FROM to_jsonb(OLD) -> key
		) sub;
	ELSIF TG_OP = 'DELETE' THEN
		v_old_values := to_jsonb(OLD);
		v_record_id := (OLD).id;
	END IF;

	INSERT INTO audit_logs (
		table_name,
		record_id,
		action,
		old_values,
		new_values,
		changed_fields
	) VALUES (
		TG_TABLE_NAME,
		v_record_id,
		TG_OP::audit_action,
		v_old_values,
		v_new_values,
		v_changed_fields
	);

	IF TG_OP = 'DELETE' THEN
		RETURN OLD;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

COMMENT ON FUNCTION fn_audit_log() IS 'Trigger function สำหรับบันทึก audit log อัตโนมัติเมื่อมีการเปลี่ยนแปลงข้อมูลในตาราง sensitive';

--bun:split

-- Trigger: members
CREATE TRIGGER trg_audit_members
	AFTER INSERT OR UPDATE OR DELETE ON members
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_members ON members IS 'บันทึก audit log ทุกการเปลี่ยนแปลงข้อมูลสมาชิก';

--bun:split

-- Trigger: payments
CREATE TRIGGER trg_audit_payments
	AFTER INSERT OR UPDATE OR DELETE ON payments
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_payments ON payments IS 'บันทึก audit log ทุกรายการชำระเงิน';

--bun:split

-- Trigger: orders
CREATE TRIGGER trg_audit_orders
	AFTER INSERT OR UPDATE OR DELETE ON orders
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_orders ON orders IS 'บันทึก audit log ทุกการเปลี่ยนแปลงออเดอร์';

--bun:split

-- Trigger: shops
CREATE TRIGGER trg_audit_shops
	AFTER INSERT OR UPDATE OR DELETE ON shops
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_shops ON shops IS 'บันทึก audit log ทุกการเปลี่ยนแปลงข้อมูลร้านค้า';

--bun:split

-- Trigger: refund_transactions
CREATE TRIGGER trg_audit_refund_transactions
	AFTER INSERT OR UPDATE OR DELETE ON refund_transactions
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_refund_transactions ON refund_transactions IS 'บันทึก audit log ทุกธุรกรรมคืนเงิน';

--bun:split

-- Trigger: settlement_batches
CREATE TRIGGER trg_audit_settlement_batches
	AFTER INSERT OR UPDATE OR DELETE ON settlement_batches
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_settlement_batches ON settlement_batches IS 'บันทึก audit log ทุก settlement batch';

--bun:split

-- Trigger: shop_wallet_transactions
CREATE TRIGGER trg_audit_shop_wallet_transactions
	AFTER INSERT OR UPDATE OR DELETE ON shop_wallet_transactions
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_shop_wallet_transactions ON shop_wallet_transactions IS 'บันทึก audit log ทุกธุรกรรมกระเป๋าเงินร้านค้า';

--bun:split

-- Trigger: kyc_verifications
CREATE TRIGGER trg_audit_kyc_verifications
	AFTER INSERT OR UPDATE OR DELETE ON kyc_verifications
	FOR EACH ROW EXECUTE FUNCTION fn_audit_log();

COMMENT ON TRIGGER trg_audit_kyc_verifications ON kyc_verifications IS 'บันทึก audit log ทุกการเปลี่ยนสถานะ KYC';
