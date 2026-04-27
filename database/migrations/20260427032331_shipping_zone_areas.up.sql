SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE shipping_zone_areas (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	shipping_zone_id uuid REFERENCES shipping_zones(id),
	province_id uuid REFERENCES provinces(id),
	district_id uuid REFERENCES districts(id),
	sub_district_id uuid REFERENCES sub_districts(id),
	zipcode_id uuid REFERENCES zipcodes(id),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	UNIQUE (shipping_zone_id, province_id, district_id, sub_district_id, zipcode_id)
);

--bun:split

COMMENT ON TABLE shipping_zone_areas IS 'ตารางกำหนดพื้นที่จริงที่อยู่ในโซนการจัดส่ง';
COMMENT ON COLUMN shipping_zone_areas.shipping_zone_id IS 'อ้างอิงโซนการจัดส่ง';
COMMENT ON COLUMN shipping_zone_areas.province_id IS 'จังหวัดในโซน';
COMMENT ON COLUMN shipping_zone_areas.district_id IS 'อำเภอหรือเขตในโซน';
COMMENT ON COLUMN shipping_zone_areas.sub_district_id IS 'ตำบลหรือแขวงในโซน';
COMMENT ON COLUMN shipping_zone_areas.zipcode_id IS 'รหัสไปรษณีย์ในโซน';
