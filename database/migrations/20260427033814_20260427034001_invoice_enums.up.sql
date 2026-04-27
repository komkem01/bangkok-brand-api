SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'invoice_status') THEN
		CREATE TYPE invoice_status AS ENUM (
			'draft',
			'issued',
			'cancelled',
			'voided'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'invoice_type') THEN
		CREATE TYPE invoice_type AS ENUM (
			'full_tax',
			'simplified',
			'receipt',
			'credit_note',
			'debit_note'
		);
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE invoice_status IS 'สถานะใบกำกับภาษี/ใบเสร็จ';
COMMENT ON TYPE invoice_type IS 'ประเภทเอกสารทางภาษี';
