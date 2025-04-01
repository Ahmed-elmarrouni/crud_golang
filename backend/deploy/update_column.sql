-- Deploy crud_golang:update_column to pg

BEGIN;

-- XXX Add DDLs here.
ALTER TABLE players ALTER COLUMN email SET NOT NULL;

COMMIT;
