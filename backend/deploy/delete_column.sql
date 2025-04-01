-- Deploy crud_golang:delete_column to pg

BEGIN;

-- XXX Add DDLs here.
ALTER TABLE players DROP COLUMN birth_date;

COMMIT;
