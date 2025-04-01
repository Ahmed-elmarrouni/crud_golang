-- Verify crud_golang:players on pg

BEGIN;

-- XXX Add verifications here.
SELECT 1 FROM players LIMIT 1;

ROLLBACK;
