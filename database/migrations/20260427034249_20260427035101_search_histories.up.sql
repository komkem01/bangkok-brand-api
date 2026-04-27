SET statement_timeout = 0;

--bun:split

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--bun:split

CREATE TABLE search_histories (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	member_id uuid REFERENCES members(id),
	session_id varchar,
	keyword varchar NOT NULL,
	result_count integer,
	platform varchar,
	created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_search_histories_member_id ON search_histories(member_id);
CREATE INDEX IF NOT EXISTS idx_search_histories_keyword ON search_histories(keyword);
CREATE INDEX IF NOT EXISTS idx_search_histories_created_at ON search_histories(created_at);

--bun:split

COMMENT ON TABLE search_histories IS 'ตารางประวัติการค้นหาสินค้า ใช้สำหรับวิเคราะห์ autocomplete และ recommendation';
COMMENT ON COLUMN search_histories.member_id IS 'สมาชิกที่ค้นหา (null = ผู้ใช้ทั่วไปไม่ล็อกอิน)';
COMMENT ON COLUMN search_histories.session_id IS 'Anonymous session ID สำหรับผู้ใช้ที่ยังไม่ล็อกอิน';
COMMENT ON COLUMN search_histories.keyword IS 'คำค้นหา';
COMMENT ON COLUMN search_histories.result_count IS 'จำนวนผลลัพธ์ที่พบในการค้นหา';
COMMENT ON COLUMN search_histories.platform IS 'ช่องทางที่ค้นหา เช่น ios, android, web';
