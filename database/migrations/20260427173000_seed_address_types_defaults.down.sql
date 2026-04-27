SET statement_timeout = 0;

--bun:split

DELETE FROM address_types
WHERE UPPER(name_en) IN ('HOME', 'WORK', 'OTHER');

--bun:split

COMMENT ON TABLE address_types IS 'ตารางเก็บประเภทที่อยู่';
COMMENT ON COLUMN address_types.name_th IS 'ชื่อประเภทที่อยู่ภาษาไทย';
COMMENT ON COLUMN address_types.name_en IS 'ชื่อประเภทที่อยู่ภาษาอังกฤษ';
COMMENT ON COLUMN address_types.is_active IS 'สถานะการใช้งานของประเภทที่อยู่';
