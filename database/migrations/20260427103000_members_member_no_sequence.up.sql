SET statement_timeout = 0;

--bun:split

CREATE SEQUENCE IF NOT EXISTS members_member_no_seq START WITH 1 INCREMENT BY 1 MINVALUE 1;

--bun:split

DO $$
DECLARE
	max_member_no bigint;
BEGIN
	SELECT COALESCE(MAX(member_no::bigint), 0)
	INTO max_member_no
	FROM members
	WHERE member_no ~ '^[0-9]+$';

	IF max_member_no > 0 THEN
		PERFORM setval('members_member_no_seq', max_member_no, true);
	ELSE
		PERFORM setval('members_member_no_seq', 1, false);
	END IF;
END
$$;

--bun:split

UPDATE members
SET member_no = LPAD(nextval('members_member_no_seq')::text, 6, '0')
WHERE member_no IS NULL OR BTRIM(member_no) = '';

--bun:split

ALTER TABLE members
	ALTER COLUMN member_no SET DEFAULT LPAD(nextval('members_member_no_seq')::text, 6, '0'),
	ALTER COLUMN member_no SET NOT NULL;

--bun:split

COMMENT ON TABLE members IS 'ตารางเก็บข้อมูลสมาชิกในระบบ';
COMMENT ON COLUMN members.member_no IS 'รหัสสมาชิกแบบ running number 6 หลัก เริ่มจาก 000001';
