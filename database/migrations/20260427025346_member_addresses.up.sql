SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE member_addresses (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	address_type_id uuid REFERENCES address_types(id),
	address_name varchar,
	recipient_name varchar,
	recipient_phone varchar,
	address_detail text,
	province_id uuid REFERENCES provinces(id),
	district_id uuid REFERENCES districts(id),
	sub_district_id uuid REFERENCES sub_districts(id),
	zipcode_id uuid REFERENCES zipcodes(id),
	is_default boolean NOT NULL DEFAULT false,
	latitude decimal(10, 8),
	longitude decimal(11, 8),
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);
