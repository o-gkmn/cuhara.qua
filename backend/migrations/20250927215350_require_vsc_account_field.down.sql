-- +migrate Down

ALTER TABLE users ALTER COLUMN vsc_account SET NULL;