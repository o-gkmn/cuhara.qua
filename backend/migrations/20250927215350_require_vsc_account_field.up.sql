-- +migrate Up

ALTER TABLE users ALTER COLUMN vsc_account SET NOT NULL;