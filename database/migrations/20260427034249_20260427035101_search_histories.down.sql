SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_search_histories_created_at;
DROP INDEX IF EXISTS idx_search_histories_keyword;
DROP INDEX IF EXISTS idx_search_histories_member_id;
DROP TABLE IF EXISTS search_histories;
