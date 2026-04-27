-- Add deleted_at column to genders table for soft delete support
ALTER TABLE genders ADD COLUMN deleted_at timestamp;

-- Create index on deleted_at for soft delete queries
CREATE INDEX idx_genders_deleted_at ON genders(deleted_at);
