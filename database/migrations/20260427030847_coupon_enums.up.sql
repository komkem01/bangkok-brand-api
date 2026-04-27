SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'coupon_type') THEN
		CREATE TYPE coupon_type AS ENUM ('fixed_amount', 'percentage', 'free_shipping');
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE coupon_type IS 'ประเภทส่วนลดของคูปอง';
