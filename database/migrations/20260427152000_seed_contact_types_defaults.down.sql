SET statement_timeout = 0;

--bun:split

DELETE FROM contact_types
WHERE UPPER(name_en) IN ('EMAIL', 'PHONE');

--bun:split

COMMENT ON TABLE contact_types IS 'ตารางเก็บประเภทช่องทางการติดต่อ';
COMMENT ON COLUMN contact_types.name_th IS 'ชื่อประเภทการติดต่อภาษาไทย';
COMMENT ON COLUMN contact_types.name_en IS 'ชื่อประเภทการติดต่อภาษาอังกฤษ';
COMMENT ON COLUMN contact_types.is_active IS 'สถานะการใช้งานของประเภทการติดต่อ';
