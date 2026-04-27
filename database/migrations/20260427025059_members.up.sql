SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'member_role') THEN
		CREATE TYPE member_role AS ENUM ('customer', 'admin', 'merchant');
	END IF;
END
$$;

--bun:split

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'member_status') THEN
		CREATE TYPE member_status AS ENUM ('active', 'inactive', 'suspended');
	END IF;
END
$$;

--bun:split

CREATE TABLE members (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	gender_id uuid REFERENCES genders(id),
	prefix_id uuid REFERENCES prefixes(id),
	email varchar UNIQUE,
	password text,
	member_no varchar UNIQUE,
	profile_image_id uuid REFERENCES storages(id),
	displayname varchar,
	firstname_th varchar,
	lastname_th varchar,
	citizen_id varchar,
	birthdate date,
	phone varchar,
	role member_role NOT NULL DEFAULT 'customer',
	status member_status NOT NULL DEFAULT 'active',
	province_id uuid REFERENCES provinces(id),
	district_id uuid REFERENCES districts(id),
	sub_district_id uuid REFERENCES sub_districts(id),
	zipcode_id uuid REFERENCES zipcodes(id),
	registerd_at timestamp,
	lasted_login timestamp,
	is_verified boolean NOT NULL DEFAULT false,
	total_points integer NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	deleted_at timestamp
);

--bun:split

COMMENT ON TABLE members IS 'ตารางเก็บข้อมูลสมาชิกในระบบ';
COMMENT ON COLUMN members.profile_image_id IS 'อ้างอิงรูปโปรไฟล์จากตาราง storage';
COMMENT ON COLUMN members.role IS 'บทบาทของสมาชิกในระบบ';
COMMENT ON COLUMN members.status IS 'สถานะของสมาชิก';
COMMENT ON COLUMN members.registerd_at IS 'วันที่และเวลาที่สมาชิกลงทะเบียน';
COMMENT ON COLUMN members.lasted_login IS 'วันที่และเวลาที่สมาชิกเข้าสู่ระบบล่าสุด';
COMMENT ON COLUMN members.is_verified IS 'สถานะการยืนยันตัวตนหรือบัญชี';
COMMENT ON COLUMN members.total_points IS 'คะแนนสะสมรวมของสมาชิก';
