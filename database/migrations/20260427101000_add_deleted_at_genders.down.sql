-- Rollback: Remove deleted_at column and index from genders table
DROP INDEX IF EXISTS idx_genders_deleted_at;
ALTER TABLE genders DROP COLUMN IF EXISTS deleted_at;
