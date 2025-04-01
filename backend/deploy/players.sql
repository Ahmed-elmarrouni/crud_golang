-- Deploy crud_golang:players to pg

BEGIN;

-- XXX Add DDLs here.
CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

COMMIT;
