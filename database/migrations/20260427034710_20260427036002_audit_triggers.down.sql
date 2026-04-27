SET statement_timeout = 0;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_kyc_verifications ON kyc_verifications;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_shop_wallet_transactions ON shop_wallet_transactions;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_settlement_batches ON settlement_batches;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_refund_transactions ON refund_transactions;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_shops ON shops;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_orders ON orders;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_payments ON payments;

--bun:split

DROP TRIGGER IF EXISTS trg_audit_members ON members;

--bun:split

DROP FUNCTION IF EXISTS fn_audit_log();
