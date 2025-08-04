-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    access_finance BOOL NOT NULL,
    access_timeregistration BOOL NOT NULL,
    administrator BOOL NOT NULL
);


-- +goose Down
DROP TABLE users;