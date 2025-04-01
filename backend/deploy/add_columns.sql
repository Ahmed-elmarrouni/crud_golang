-- Deploy crud_golang:add_columns to pg

BEGIN;

-- XXX Add DDLs here.
ALTER TABLE players ADD COLUMN email TEXT;
ALTER TABLE players ADD COLUMN birth_date DATE;

COMMIT;
