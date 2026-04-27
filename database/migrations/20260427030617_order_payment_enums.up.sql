SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
		CREATE TYPE order_status AS ENUM (
			'pending_payment',
			'awaiting_confirmation',
			'processing',
			'shipped',
			'delivered',
			'cancelled',
			'refunded'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_method') THEN
		CREATE TYPE payment_method AS ENUM (
			'bank_transfer',
			'credit_card',
			'prompt_pay',
			'cash_on_delivery'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status') THEN
		CREATE TYPE payment_status AS ENUM (
			'pending',
			'completed',
			'failed',
			'refunded'
		);
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE order_status IS 'สถานะของคำสั่งซื้อ';
COMMENT ON TYPE payment_method IS 'ช่องทางการชำระเงิน';
COMMENT ON TYPE payment_status IS 'สถานะการชำระเงิน';
