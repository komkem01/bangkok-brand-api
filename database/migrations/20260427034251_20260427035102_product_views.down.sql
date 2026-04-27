SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_product_views_viewed_at;
DROP INDEX IF EXISTS idx_product_views_product_id;
DROP INDEX IF EXISTS idx_product_views_member_id;
DROP TABLE IF EXISTS product_views;
