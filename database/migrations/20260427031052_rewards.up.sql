SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE rewards (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name_th varchar,
	name_en varchar,
	description text,
	type reward_type,
	points_required integer,
	image_id uuid REFERENCES storages(id),
	stock_quantity integer NOT NULL DEFAULT 0,
	start_date timestamp,
	end_date timestamp,
	is_active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp
);

--bun:split

COMMENT ON TABLE rewards IS 'ตารางเก็บข้อมูลของรางวัลสำหรับแลกคะแนน';
COMMENT ON COLUMN rewards.points_required IS 'จำนวนคะแนนที่ต้องใช้แลก';
COMMENT ON COLUMN rewards.image_id IS 'รูปภาพของรางวัล';
COMMENT ON COLUMN rewards.stock_quantity IS 'จำนวนของรางวัลที่มีให้แลก';
