SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE storages (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	file_name varchar,
	file_path text,
	file_extension varchar(10),
	file_size integer,
	mime_type varchar,
	provider varchar NOT NULL DEFAULT 'local',
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	deleted_at timestamp
);

--bun:split

COMMENT ON TABLE storages IS 'ตารางกลางสำหรับเก็บข้อมูลไฟล์และรูปภาพทั้งหมดในระบบ';
COMMENT ON COLUMN storages.file_name IS 'ชื่อไฟล์ต้นฉบับ';
COMMENT ON COLUMN storages.file_path IS 'ที่อยู่ไฟล์ในระบบ/Cloud Storage';
COMMENT ON COLUMN storages.file_extension IS 'นามสกุลไฟล์ เช่น jpg, png, pdf';
COMMENT ON COLUMN storages.file_size IS 'ขนาดไฟล์ (bytes)';
COMMENT ON COLUMN storages.mime_type IS 'ประเภทไฟล์ เช่น image/jpeg';
COMMENT ON COLUMN storages.provider IS 'ที่เก็บไฟล์ เช่น local, s3, cloudinary';
