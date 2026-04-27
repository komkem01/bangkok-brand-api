SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_audit_logs_created_at;
DROP INDEX IF EXISTS idx_audit_logs_actor_id;
DROP INDEX IF EXISTS idx_audit_logs_table_record;
DROP TABLE IF EXISTS audit_logs;

--bun:split

DROP TYPE IF EXISTS audit_action;
