SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_notifications_member_unread;
DROP INDEX IF EXISTS idx_notifications_member_id;
DROP TABLE IF EXISTS notifications;
