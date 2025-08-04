-- +goose Up
ALTER TABLE timeregistration
ALTER COLUMN timestamp TYPE DATE,
ALTER COLUMN date_activity TYPE DATE;

ALTER TABLE finances
ALTER COLUMN timestamp TYPE DATE,
ALTER COLUMN date_transaction TYPE DATE;

-- +goose Down
ALTER TABLE timeregistration
ALTER COLUMN timestamp TYPE TIMESTAMP,
ALTER COLUMN date_activity TYPE TIMESTAMP;

ALTER TABLE finances
ALTER COLUMN timestamp TYPE TIMESTAMP,
ALTER COLUMN date_transaction TYPE TIMESTAMP;