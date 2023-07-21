BEGIN;

ALTER TABLE users ADD COLUMN last_activity timestamptz NOT NULL DEFAULT (now());

COMMIT;