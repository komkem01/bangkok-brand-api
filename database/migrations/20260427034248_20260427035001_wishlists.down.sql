SET statement_timeout = 0;

--bun:split

DROP INDEX IF EXISTS idx_wishlists_product_id;
DROP INDEX IF EXISTS idx_wishlists_member_id;
DROP TABLE IF EXISTS wishlists;
