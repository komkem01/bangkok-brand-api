SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'point_transaction_type') THEN
		CREATE TYPE point_transaction_type AS ENUM (
			'earn',
			'redeem',
			'expired',
			'refund_adjust',
			'admin_adjust'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reward_type') THEN
		CREATE TYPE reward_type AS ENUM ('discount_voucher', 'physical_product', 'service_gift');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'redemption_status') THEN
		CREATE TYPE redemption_status AS ENUM ('pending', 'completed', 'cancelled');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'message_type') THEN
		CREATE TYPE message_type AS ENUM ('text', 'image', 'file', 'system');
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE point_transaction_type IS 'ประเภทรายการเคลื่อนไหวคะแนนสะสม';
COMMENT ON TYPE reward_type IS 'ประเภทของรางวัล';
COMMENT ON TYPE redemption_status IS 'สถานะการแลกรางวัล';
COMMENT ON TYPE message_type IS 'ประเภทข้อความในระบบแชท';
