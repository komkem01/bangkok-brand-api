SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_webhook_events_pending_retry;
DROP INDEX IF EXISTS idx_webhook_events_status;
DROP INDEX IF EXISTS idx_webhook_events_shop_id;
DROP TABLE IF EXISTS webhook_events;

--bun:split

DROP TYPE IF EXISTS webhook_event_status;
