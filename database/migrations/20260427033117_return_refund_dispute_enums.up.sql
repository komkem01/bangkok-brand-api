SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'return_request_status') THEN
		CREATE TYPE return_request_status AS ENUM (
			'requested',
			'approved',
			'rejected',
			'received',
			'refunded',
			'cancelled'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'refund_transaction_status') THEN
		CREATE TYPE refund_transaction_status AS ENUM (
			'pending',
			'processing',
			'completed',
			'failed',
			'cancelled'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'dispute_case_status') THEN
		CREATE TYPE dispute_case_status AS ENUM ('open', 'under_review', 'resolved', 'rejected', 'closed');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'dispute_message_party') THEN
		CREATE TYPE dispute_message_party AS ENUM ('customer', 'merchant', 'admin', 'system');
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE return_request_status IS 'สถานะคำขอคืนสินค้า';
COMMENT ON TYPE refund_transaction_status IS 'สถานะการคืนเงิน';
COMMENT ON TYPE dispute_case_status IS 'สถานะเคสข้อพิพาท';
COMMENT ON TYPE dispute_message_party IS 'ฝั่งผู้ส่งข้อความในข้อพิพาท';
