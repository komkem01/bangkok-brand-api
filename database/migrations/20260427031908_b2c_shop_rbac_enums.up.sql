SET statement_timeout = 0;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'shop_status') THEN
		CREATE TYPE shop_status AS ENUM ('pending_kyc', 'active', 'suspended', 'closed');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'shop_member_role') THEN
		CREATE TYPE shop_member_role AS ENUM ('owner', 'manager', 'staff');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'admin_action_type') THEN
		CREATE TYPE admin_action_type AS ENUM (
			'create',
			'update',
			'delete',
			'approve',
			'reject',
			'suspend',
			'unsuspend',
			'refund',
			'settlement'
		);
	END IF;
END
$$;

--bun:split

COMMENT ON TYPE shop_status IS 'สถานะการดำเนินงานของร้านค้า';
COMMENT ON TYPE shop_member_role IS 'สิทธิ์ของสมาชิกภายในร้านค้า';
COMMENT ON TYPE admin_action_type IS 'ประเภทกิจกรรมที่แอดมินดำเนินการในระบบ';
