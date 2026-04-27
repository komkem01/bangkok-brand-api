SET statement_timeout = 0;

--bun:split

COMMENT ON TABLE contact_types IS NULL;
COMMENT ON COLUMN contact_types.name_th IS NULL;
COMMENT ON COLUMN contact_types.name_en IS NULL;
COMMENT ON COLUMN contact_types.is_active IS NULL;

--bun:split

COMMENT ON TABLE member_contacts IS NULL;
COMMENT ON COLUMN member_contacts.member_id IS NULL;
COMMENT ON COLUMN member_contacts.contact_type_id IS NULL;
COMMENT ON COLUMN member_contacts.value IS NULL;
COMMENT ON COLUMN member_contacts.is_primary IS NULL;
COMMENT ON COLUMN member_contacts.is_verified IS NULL;

--bun:split

COMMENT ON TABLE address_types IS NULL;
COMMENT ON COLUMN address_types.name_th IS NULL;
COMMENT ON COLUMN address_types.name_en IS NULL;
COMMENT ON COLUMN address_types.is_active IS NULL;

--bun:split

COMMENT ON TABLE member_addresses IS NULL;
COMMENT ON COLUMN member_addresses.member_id IS NULL;
COMMENT ON COLUMN member_addresses.address_type_id IS NULL;
COMMENT ON COLUMN member_addresses.address_name IS NULL;
COMMENT ON COLUMN member_addresses.recipient_name IS NULL;
COMMENT ON COLUMN member_addresses.recipient_phone IS NULL;
COMMENT ON COLUMN member_addresses.address_detail IS NULL;
COMMENT ON COLUMN member_addresses.province_id IS NULL;
COMMENT ON COLUMN member_addresses.district_id IS NULL;
COMMENT ON COLUMN member_addresses.sub_district_id IS NULL;
COMMENT ON COLUMN member_addresses.zipcode_id IS NULL;
COMMENT ON COLUMN member_addresses.is_default IS NULL;
COMMENT ON COLUMN member_addresses.latitude IS NULL;
COMMENT ON COLUMN member_addresses.longitude IS NULL;
