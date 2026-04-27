SET statement_timeout = 0;

--bun:split

INSERT INTO address_types (name_th, name_en, is_active)
SELECT 'บ้าน', 'HOME', true
WHERE NOT EXISTS (
	SELECT 1 FROM address_types WHERE UPPER(name_en) = 'HOME'
);

--bun:split

INSERT INTO address_types (name_th, name_en, is_active)
SELECT 'ที่ทำงาน', 'WORK', true
WHERE NOT EXISTS (
	SELECT 1 FROM address_types WHERE UPPER(name_en) = 'WORK'
);

--bun:split

INSERT INTO address_types (name_th, name_en, is_active)
SELECT 'อื่นๆ', 'OTHER', true
WHERE NOT EXISTS (
	SELECT 1 FROM address_types WHERE UPPER(name_en) = 'OTHER'
);

--bun:split

COMMENT ON TABLE address_types IS 'ตารางเก็บประเภทที่อยู่';
COMMENT ON COLUMN address_types.name_th IS 'ชื่อประเภทที่อยู่ภาษาไทย';
COMMENT ON COLUMN address_types.name_en IS 'ชื่อประเภทที่อยู่ภาษาอังกฤษ';
COMMENT ON COLUMN address_types.is_active IS 'สถานะการใช้งานของประเภทที่อยู่';
