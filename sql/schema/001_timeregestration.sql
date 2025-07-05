-- +goose Up
CREATE TABLE tijdregistratie(
    id UUID PRIMARY KEY,
    tijd TIMESTAMP NOT NULL,
    lengte_minuten integer NOT NULL,
    beschrijving TEXT NOT NULL,
    catagorie TEXT NOT NULL
);

CREATE TABLE financien(
    id UUID PRIMARY KEY,
    tijdregistratie TIMESTAMP NOT NULL,
    bedrag_cent INTEGER NOT NULL,
    type TEXT NOT NULL,
    beschrijving TEXT NOT NULL,
    catagorie TEXT NOT NULL
);

-- +goose Down
DROP TABLE financien
DROP TABLE tijdregistratie;