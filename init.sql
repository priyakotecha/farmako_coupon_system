CREATE DATABASE IF NOT EXISTS coupon_db;
\c coupon_db;

CREATE TABLE IF NOT EXISTS coupons (
    code TEXT PRIMARY KEY,
    expiry_date TIMESTAMP,
    usage_type TEXT,
    min_order_value FLOAT,
    valid_from TIMESTAMP,
    valid_until TIMESTAMP,
    terms_and_conditions TEXT,
    discount_type TEXT,
    discount_value FLOAT,
    max_usage_per_user INT
);

INSERT INTO coupons (code, expiry_date, usage_type, min_order_value, valid_from, valid_until, terms_and_conditions, discount_type, discount_value, max_usage_per_user)
VALUES
('SAVE20', '2025-12-31T23:59:59', 'multi_use', 500, '2025-01-01T00:00:00', '2025-12-31T23:59:59', 'Save 20 on your order', 'inventory', 20, 10)
ON CONFLICT (code) DO NOTHING;