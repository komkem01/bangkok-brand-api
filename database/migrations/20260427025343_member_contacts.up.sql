SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE member_contacts (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	contact_type_id uuid REFERENCES contact_types(id),
	value varchar,
	is_primary boolean NOT NULL DEFAULT false,
	is_verified boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);
