-- Rollback: Remove deleted_at column and index from prefixes table
DROP INDEX IF EXISTS idx_prefixes_deleted_at;
ALTER TABLE prefixes DROP COLUMN IF EXISTS deleted_at;
