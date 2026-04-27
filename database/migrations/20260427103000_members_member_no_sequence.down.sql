SET statement_timeout = 0;

--bun:split

ALTER TABLE members
	ALTER COLUMN member_no DROP NOT NULL,
	ALTER COLUMN member_no DROP DEFAULT;

--bun:split

DROP SEQUENCE IF EXISTS members_member_no_seq;

--bun:split

COMMENT ON TABLE members IS 'ตารางเก็บข้อมูลสมาชิกในระบบ';
COMMENT ON COLUMN members.member_no IS NULL;
