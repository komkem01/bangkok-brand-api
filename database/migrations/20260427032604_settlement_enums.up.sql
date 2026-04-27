SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'settlement_status') THEN
		CREATE TYPE settlement_status AS ENUM (
			'pending',
			'processing',
			'approved',
			'paid',
			'failed',
			'cancelled'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'wallet_transaction_type') THEN
		CREATE TYPE wallet_transaction_type AS ENUM (
			'order_income',
			'shipping_income',
			'platform_fee',
			'refund_adjust',
			'admin_adjust',
			'payout',
			'reserve_hold',
			'reserve_release'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'settlement_item_status') THEN
		CREATE TYPE settlement_item_status AS ENUM ('pending', 'included', 'skipped', 'paid');
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE settlement_status IS 'สถานะรอบการโอนเงินให้ร้านค้า';
COMMENT ON TYPE wallet_transaction_type IS 'ประเภทรายการในกระเป๋าเงินร้านค้า';
COMMENT ON TYPE settlement_item_status IS 'สถานะของรายการย่อยในรอบโอนเงิน';
