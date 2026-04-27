-- Add deleted_at column to prefixes table for soft delete support
ALTER TABLE prefixes ADD COLUMN deleted_at timestamp;

-- Create index on deleted_at for soft delete queries
CREATE INDEX idx_prefixes_deleted_at ON prefixes(deleted_at);
