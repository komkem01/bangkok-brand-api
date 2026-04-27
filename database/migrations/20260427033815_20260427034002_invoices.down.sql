SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_invoices_shop_id;
DROP INDEX IF EXISTS idx_invoices_member_id;
DROP INDEX IF EXISTS idx_invoices_order_id;
DROP TABLE IF EXISTS invoices;
