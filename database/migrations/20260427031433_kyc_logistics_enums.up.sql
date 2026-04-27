SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kyc_status') THEN
		CREATE TYPE kyc_status AS ENUM (
			'pending',
			'in_review',
			'approved',
			'rejected',
			'resubmitted',
			'suspended'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kyc_entity_type') THEN
		CREATE TYPE kyc_entity_type AS ENUM ('merchant', 'community_enterprise');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kyc_document_type') THEN
		CREATE TYPE kyc_document_type AS ENUM (
			'national_id',
			'business_registration',
			'tax_document',
			'bank_book',
			'address_proof',
			'other'
		);
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'shipping_method_type') THEN
		CREATE TYPE shipping_method_type AS ENUM ('standard', 'express', 'same_day', 'pickup');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'shipment_status') THEN
		CREATE TYPE shipment_status AS ENUM (
			'pending_pickup',
			'picked_up',
			'in_transit',
			'out_for_delivery',
			'delivered',
			'failed',
			'returned',
			'cancelled'
		);
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE kyc_status IS 'สถานะการตรวจสอบ KYC';
COMMENT ON TYPE kyc_entity_type IS 'ประเภทผู้ยื่น KYC';
COMMENT ON TYPE kyc_document_type IS 'ประเภทเอกสาร KYC';
COMMENT ON TYPE shipping_method_type IS 'ประเภทช่องทางการจัดส่ง';
COMMENT ON TYPE shipment_status IS 'สถานะการขนส่งพัสดุ';
