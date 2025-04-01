-- Revert crud_golang:players from pg

BEGIN;

-- XXX Add DDLs here.
DROP TABLE players;

COMMIT;
