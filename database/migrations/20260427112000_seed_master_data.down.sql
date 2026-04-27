SET statement_timeout = 0;

--bun:split

-- Roll back seeded zipcode data first (child -> parent).
DELETE FROM zipcodes z
USING sub_districts s, districts d, provinces p
WHERE z.sub_district_id = s.id
  AND s.district_id = d.id
  AND d.province_id = p.id
  AND p.name IN ('กรุงเทพมหานคร', 'เชียงใหม่', 'ชลบุรี')
  AND (
      (d.name = 'เขตบางรัก' AND s.name IN ('สีลม', 'สุริยวงศ์') AND z.name = '10500') OR
      (d.name = 'เขตจตุจักร' AND s.name IN ('จตุจักร', 'ลาดยาว') AND z.name = '10900') OR
      (d.name = 'เมืองเชียงใหม่' AND s.name IN ('ศรีภูมิ', 'พระสิงห์') AND z.name = '50200') OR
      (d.name = 'สันทราย' AND s.name IN ('หนองจ๊อม', 'สันทรายน้อย') AND z.name = '50210') OR
      (d.name = 'เมืองชลบุรี' AND s.name IN ('บ้านสวน', 'บางปลาสร้อย') AND z.name = '20000') OR
      (d.name = 'บางละมุง' AND s.name IN ('นาเกลือ', 'หนองปรือ') AND z.name = '20150')
  );

--bun:split

DELETE FROM sub_districts s
USING districts d, provinces p
WHERE s.district_id = d.id
  AND d.province_id = p.id
  AND p.name IN ('กรุงเทพมหานคร', 'เชียงใหม่', 'ชลบุรี')
  AND (
      (d.name = 'เขตบางรัก' AND s.name IN ('สีลม', 'สุริยวงศ์')) OR
      (d.name = 'เขตจตุจักร' AND s.name IN ('จตุจักร', 'ลาดยาว')) OR
      (d.name = 'เมืองเชียงใหม่' AND s.name IN ('ศรีภูมิ', 'พระสิงห์')) OR
      (d.name = 'สันทราย' AND s.name IN ('หนองจ๊อม', 'สันทรายน้อย')) OR
      (d.name = 'เมืองชลบุรี' AND s.name IN ('บ้านสวน', 'บางปลาสร้อย')) OR
      (d.name = 'บางละมุง' AND s.name IN ('นาเกลือ', 'หนองปรือ'))
  );

--bun:split

DELETE FROM districts d
USING provinces p
WHERE d.province_id = p.id
  AND (
      (p.name = 'กรุงเทพมหานคร' AND d.name IN ('เขตบางรัก', 'เขตจตุจักร')) OR
      (p.name = 'เชียงใหม่' AND d.name IN ('เมืองเชียงใหม่', 'สันทราย')) OR
      (p.name = 'ชลบุรี' AND d.name IN ('เมืองชลบุรี', 'บางละมุง'))
  );

--bun:split

DELETE FROM provinces
WHERE name IN ('กรุงเทพมหานคร', 'เชียงใหม่', 'ชลบุรี');

--bun:split

DELETE FROM prefixes
WHERE name_th IN ('นาย', 'นาง', 'นางสาว', 'เด็กชาย', 'เด็กหญิง', 'คุณ', 'ดร.', 'ผศ.');

--bun:split

DELETE FROM genders
WHERE name_th IN ('ชาย', 'หญิง', 'ไม่ระบุ');
