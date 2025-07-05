-- +goose Up
CREATE TABLE timeregistration(
    id UUID PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    date_activity TIMESTAMP NOT NULL,
    length_minutes integer NOT NULL,
    description TEXT NOT NULL,
    catagory TEXT NOT NULL
);

CREATE TABLE finances(
    id UUID PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    date_transaction TIMESTAMP NOT NULL,
    ammount_cent INTEGER NOT NULL,
    type TEXT NOT NULL,
    description TEXT NOT NULL,
    catagory TEXT NOT NULL
);

-- +goose Down
DROP TABLE finances;
DROP TABLE timeregistration;