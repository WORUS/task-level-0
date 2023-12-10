CREATE TABLE IF NOT EXISTS orders
(
    order_uid UUID UNIQUE,
    order_json jsonb
);