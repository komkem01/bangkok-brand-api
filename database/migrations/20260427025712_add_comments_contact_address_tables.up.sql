SET statement_timeout = 0;

--bun:split

COMMENT ON TABLE contact_types IS 'ตารางเก็บประเภทช่องทางการติดต่อ';
COMMENT ON COLUMN contact_types.name_th IS 'ชื่อประเภทการติดต่อภาษาไทย';
COMMENT ON COLUMN contact_types.name_en IS 'ชื่อประเภทการติดต่อภาษาอังกฤษ';
COMMENT ON COLUMN contact_types.is_active IS 'สถานะการใช้งานของประเภทการติดต่อ';

--bun:split

COMMENT ON TABLE member_contacts IS 'ตารางเก็บข้อมูลการติดต่อของสมาชิก';
COMMENT ON COLUMN member_contacts.member_id IS 'อ้างอิงสมาชิกเจ้าของข้อมูลการติดต่อ';
COMMENT ON COLUMN member_contacts.contact_type_id IS 'อ้างอิงประเภทช่องทางการติดต่อ';
COMMENT ON COLUMN member_contacts.value IS 'ค่าข้อมูลการติดต่อ เช่น อีเมลหรือไอดีโซเชียล';
COMMENT ON COLUMN member_contacts.is_primary IS 'ระบุว่าเป็นช่องทางติดต่อหลักของสมาชิก';
COMMENT ON COLUMN member_contacts.is_verified IS 'สถานะยืนยันข้อมูลการติดต่อ';

--bun:split

COMMENT ON TABLE address_types IS 'ตารางเก็บประเภทที่อยู่';
COMMENT ON COLUMN address_types.name_th IS 'ชื่อประเภทที่อยู่ภาษาไทย';
COMMENT ON COLUMN address_types.name_en IS 'ชื่อประเภทที่อยู่ภาษาอังกฤษ';
COMMENT ON COLUMN address_types.is_active IS 'สถานะการใช้งานของประเภทที่อยู่';

--bun:split

COMMENT ON TABLE member_addresses IS 'ตารางเก็บข้อมูลที่อยู่ของสมาชิก';
COMMENT ON COLUMN member_addresses.member_id IS 'อ้างอิงสมาชิกเจ้าของที่อยู่';
COMMENT ON COLUMN member_addresses.address_type_id IS 'อ้างอิงประเภทที่อยู่';
COMMENT ON COLUMN member_addresses.address_name IS 'ชื่อเรียกที่อยู่ เช่น บ้านหรือที่ทำงาน';
COMMENT ON COLUMN member_addresses.recipient_name IS 'ชื่อผู้รับสินค้า/ผู้ติดต่อ';
COMMENT ON COLUMN member_addresses.recipient_phone IS 'เบอร์โทรศัพท์ผู้รับ';
COMMENT ON COLUMN member_addresses.address_detail IS 'รายละเอียดที่อยู่';
COMMENT ON COLUMN member_addresses.province_id IS 'อ้างอิงจังหวัดของที่อยู่';
COMMENT ON COLUMN member_addresses.district_id IS 'อ้างอิงอำเภอหรือเขตของที่อยู่';
COMMENT ON COLUMN member_addresses.sub_district_id IS 'อ้างอิงตำบลหรือแขวงของที่อยู่';
COMMENT ON COLUMN member_addresses.zipcode_id IS 'อ้างอิงรหัสไปรษณีย์ของที่อยู่';
COMMENT ON COLUMN member_addresses.is_default IS 'ระบุว่าเป็นที่อยู่เริ่มต้นของสมาชิก';
COMMENT ON COLUMN member_addresses.latitude IS 'ค่าพิกัดละติจูดของที่อยู่';
COMMENT ON COLUMN member_addresses.longitude IS 'ค่าพิกัดลองจิจูดของที่อยู่';
