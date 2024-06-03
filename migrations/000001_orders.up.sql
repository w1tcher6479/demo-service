CREATE TABLE orders (
    order_uid VARCHAR PRIMARY KEY,
    data jsonb UNIQUE
);