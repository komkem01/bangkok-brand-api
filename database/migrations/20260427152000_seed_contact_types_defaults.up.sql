SET statement_timeout = 0;

--bun:split

INSERT INTO contact_types (name_th, name_en, is_active)
SELECT 'อีเมล', 'EMAIL', true
WHERE NOT EXISTS (
	SELECT 1 FROM contact_types WHERE UPPER(name_en) = 'EMAIL'
);

--bun:split

INSERT INTO contact_types (name_th, name_en, is_active)
SELECT 'เบอร์โทรศัพท์', 'PHONE', true
WHERE NOT EXISTS (
	SELECT 1 FROM contact_types WHERE UPPER(name_en) = 'PHONE'
);

--bun:split

COMMENT ON TABLE contact_types IS 'ตารางเก็บประเภทช่องทางการติดต่อ';
COMMENT ON COLUMN contact_types.name_th IS 'ชื่อประเภทการติดต่อภาษาไทย';
COMMENT ON COLUMN contact_types.name_en IS 'ชื่อประเภทการติดต่อภาษาอังกฤษ';
COMMENT ON COLUMN contact_types.is_active IS 'สถานะการใช้งานของประเภทการติดต่อ';
