CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE customer (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(200) NOT NULL,
    name VARCHAR(200) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

SELECT * FROM customer;
SELECT name FROM customer WHERE code='M-0001';
INSERT INTO customer (code, name)
VALUES ('M-0001', 'Naruto Uzumaki'),
       ('M-0002', 'Sasuke Uchiha'),
       ('M-0003', 'Sakura Haruno');
